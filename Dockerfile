FROM golang:1.5
MAINTAINER zhiqing.yang <zhiqing.yang.f@gmail.com>
ENV GOBIN /go/bin
COPY . /go
WORKDIR /go
ENV GOPATH /go:/go/.godeps
ENTRYPOINT /go/bin/ipa_service
EXPOSE 50003
