# Build stage
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache curl libstdc++ libgcc

WORKDIR /app

# Download Tailwind CSS standalone CLI
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64-musl \
    && mv tailwindcss-linux-x64-musl tailwindcss \
    && chmod +x tailwindcss

# Copy dependency files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build Tailwind CSS
RUN ./tailwindcss -i static/css/input.css -o static/css/output.css --minify

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o packstring ./cmd/server

# Runtime stage
FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/packstring .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/data ./data

ENV PORT=80
EXPOSE 80

CMD ["./packstring"]
