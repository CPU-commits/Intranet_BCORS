FROM golang:1.19.1-alpine3.16

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

RUN go mod download

ENTRYPOINT CompileDaemon -polling --build="go build main.go" --command=./main
EXPOSE 3333