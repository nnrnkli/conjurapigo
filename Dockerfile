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
    mkdir -p /summon-conjur/output

WORKDIR /summon-conjur

COPY main.go ./
RUN go mod download

COPY . .
RUN go build -o summon-conjur cmd/main.go

EXPOSE 8080