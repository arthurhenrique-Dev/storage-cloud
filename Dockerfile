# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and compile
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server .

# Runtime Stage (Imagem leve Alpine)
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
