FROM golang:1.23-alpine

COPY go.mod client/main.go /app/

WORKDIR /app

RUN ls -la . && \
    go install . && \
    which loki

ENTRYPOINT ["/go/bin/loki"]
