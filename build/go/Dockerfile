FROM golang:1.19.4

RUN apt-get update

# install golangci-lint.
# binary will be $(go env GOPATH)/bin/golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

# JSTをローカルタイムに
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /go/src/github.com/shimabox/example-s3-to-lambda-go