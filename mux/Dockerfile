FROM golang:1.19-alpine

RUN apk add --no-cache git
COPY . /go/mux

WORKDIR /go/mux
ENV GOPATH=

RUN set -ex && \
    apk add --no-cache gcc musl-dev
RUN set -ex && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/cc1obj && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/lto1 && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/lto-wrapper && \
    rm -f /usr/bin/x86_64-alpine-linux-musl-gcj

RUN go get . && go build -o /go/mux/camping-server

FROM alpine:3.13
RUN apk add --no-cache bash ca-certificates
RUN apk add --no-cache make build-base

RUN mkdir -p /usr/apps/camping-server
WORKDIR /usr/apps/camping-server
COPY --from=0 /go/mux/camping-server .


ENTRYPOINT ./camping-server