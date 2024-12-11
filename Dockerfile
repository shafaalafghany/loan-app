# Stage 1: Build stage
FROM golang:1.23.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o loan-app .

# Stage 2: Runtime stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/loan-app .
EXPOSE 3000
CMD ["./loan-app"]
