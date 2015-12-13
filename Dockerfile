FROM golang:alpine

COPY main.go /app/
CMD go run /app/main.go
