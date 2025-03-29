FROM golang:alpine AS builder
WORKDIR /app
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest
COPY . .
CMD ["air", "-c", ".air.toml"]