FROM golang:1.24.1-alpine AS builder
WORKDIR /opt

RUN apk add --no-cache git

# NOTE: use mod cache
COPY go.mod go.sum ./
RUN go mod download

# NOTE: build cache
COPY . src
RUN go build -o app ./src

# 実行ステージ
FROM alpine:latest
COPY --from=builder /opt/app .
EXPOSE 8080
CMD ["./app"]
