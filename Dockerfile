FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .
RUN go mod download
RUN go build -ldflags "-s -w" -o meme .

# Build a small image
FROM alpine
WORKDIR /dist
COPY --from=builder /build/meme .
COPY --from=builder /build/assets ./assets
ENTRYPOINT ["/dist/meme", "-dir=/dist/assets/lukashenko"]
