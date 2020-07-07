FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .
RUN go mod download
RUN go build -ldflags "-s -w" -o meme .

FROM alpine
WORKDIR /dist
COPY --from=builder /build/meme .
COPY --from=builder /build/assets ./assets
COPY --from=builder /build/assets.json ./assets.json
ENTRYPOINT ["/dist/meme"]
