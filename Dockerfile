ARG GO_VERSION=1.16.3

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY . .
RUN go mod download
RUN go get github.com/jessevdk/go-assets-builder
RUN go get github.com/GeertJohan/go.rice
RUN go get github.com/GeertJohan/go.rice/rice

WORKDIR /api/src
RUN go-assets-builder views -o assets.go
RUN rice embed-go -i .
RUN go build -o ../build/app

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/build/app .

EXPOSE 8080

ENTRYPOINT ["./app"]