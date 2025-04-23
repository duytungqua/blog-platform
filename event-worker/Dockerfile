# Dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o event-worker
CMD ["./event-worker"]