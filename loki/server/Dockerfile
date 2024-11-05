FROM golang:1.23-alpine

RUN apk add curl 

COPY go.mod server/main.go /app/

WORKDIR /app

RUN ls -la . && \
    go install . && \
    which loki

ENTRYPOINT ["/go/bin/loki"]
