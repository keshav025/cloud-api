FROM --platform=linux/amd64 golang:1.21.7-alpine3.19 AS builder
MAINTAINER kkeshav025@gmail.com
ENV GOBIN ${GOPATH}/bin
ENV PATH $PATH:$GOBIN:/usr/local/include/bin:$GOROOT/bin
RUN apk update  && apk add unzip tree
WORKDIR /go/src/cloud-api/
COPY . .
# COPY backend/commons ./backend/commons

RUN export CGO_ENABLED=0; GOOS=linux GOARCH=amd64 go build ./backend/vnet-svc

FROM --platform=linux/amd64 alpine:3.19
WORKDIR /usr/app
COPY --from=builder /go/src/cloud-api/vnet-svc .
ENV PATH $PATH:/usr/app
CMD ["vnet-svc"]