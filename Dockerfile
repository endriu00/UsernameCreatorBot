# Start from image golang
FROM golang:1.16-alpine AS build

# use /bot as working directory for the build
WORKDIR /bot

# copy each go configuration file in workdir
# and download the modules contained in go.mod
COPY go.* ./
RUN go mod download

# copy each of the directories in workdir
COPY cmd cmd
COPY service service 
COPY names names

# mkdir of the binary directory
RUN mkdir /app

# change workdir to cmd in order to build the code
WORKDIR /bot/cmd/
RUN go build -o /app/ ./

# start from the alpine image to build in a fresh environment
FROM alpine:latest

# set /app as workdir
WORKDIR /app/

# copy the output of the previous build to workdir
COPY --from=build /app/* ./

# define the command to run when running the container
ENTRYPOINT [ "/app/cmd" ]
