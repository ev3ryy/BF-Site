FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o BF ./src/
    
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/BF .
EXPOSE 8080
CMD ["./BF"]