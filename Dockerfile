FROM golang:1.21.2-alpine

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY ./ ./

ENTRYPOINT CompileDaemon -build="go build -o bin/caffeineBlues /app/cmd/caffeineBlues" -command="./bin/caffeineBlues" -build-dir=/app
