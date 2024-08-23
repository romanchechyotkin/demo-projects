FROM golang:alpine as builder
WORKDIR /app
COPY cert.pem key.pem ./
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bin main.go

FROM alpine 
WORKDIR /app
COPY --from=builder /app/cert.pem .
COPY --from=builder /app/key.pem .
COPY --from=builder /app/bin .
EXPOSE 5000
CMD ["/app/bin"]