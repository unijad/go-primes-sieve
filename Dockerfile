ARG GO_VERSION=1.16.3

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY . .
RUN go get github.com/jessevdk/go-assets-builder
RUN go mod tidy
RUN go mod download

WORKDIR /api/src
RUN go-assets-builder views -o assets.go
RUN go build -o ../build/app

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/build/app .

EXPOSE 8080

ENTRYPOINT ["./app"]