# builder
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=~/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=~/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o server server.go

# output image for app
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ENV GIN_MODE=release

CMD ["./server"]
