FROM golang:1.15-alpine

RUN apk add --no-cache bash \
                       build-base \
                       curl \
                       git \
                       jq \
                       less && \
    go get -u github.com/jstemmer/go-junit-report && \
    go get -u github.com/smartystreets/goconvey && \
    go get -u github.com/axw/gocov/gocov && \
    go get -u github.com/AlekSi/gocov-xml && \
    mkdir -p /conjurapigo/output

WORKDIR /conjurapigo

COPY main.go go.mod ./
RUN go mod download

COPY . .
RUN go build -o conjurapigo cmd/main.go

EXPOSE 8080