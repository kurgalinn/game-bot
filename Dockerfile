FROM golang:1.18 AS build

WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o bot /build/cmd/bot

FROM debian:buster-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/* \ --no-preserve-root

WORKDIR /

COPY --from=build /build /
