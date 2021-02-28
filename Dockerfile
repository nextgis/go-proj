ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine

RUN apk --no-cache --update upgrade && apk add alpine-sdk git proj proj-dev && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY . ${GOPATH}/src/gitlab.com/nextgis/go-proj/
WORKDIR ${GOPATH}/src/gitlab.com/nextgis/go-proj/
RUN go mod download && go test -v
