FROM golang:1.15.6
COPY . /go/src/github.com/brijeshshah13/url-shortener
WORKDIR /go/src/github.com/brijeshshah13/url-shortener
RUN go install -ldflags="-s -w" ./cmd/...
