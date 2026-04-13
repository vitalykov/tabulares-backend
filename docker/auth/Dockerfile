# Building
FROM golang:1.25 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build -v \
    -o /app/server \
    ./cmd/auth/main.go

# Serving
FROM alpine:3.23

EXPOSE 8081

RUN addgroup -g 1000 -S tabulares && \
adduser -u 1000 -G tabulares -S -H juggler

RUN mkdir -p /app /app/tmp && \
    chown -R juggler:tabulares /app && \
    chmod 755 /app

WORKDIR /app

COPY --from=builder --chown=juggler:tabulares /app/server /app/server

USER juggler

ENTRYPOINT [ "/app/server" ]
