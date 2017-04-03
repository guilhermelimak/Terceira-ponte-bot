FROM golang:latest

RUN mkdir $GOPATH/src/terceirapontebot

COPY . $GOPATH/src/terceirapontebot

WORKDIR $GOPATH/src/terceirapontebot

RUN go get
RUN go build

ENTRYPOINT $GOPATH/src/terceirapontebot/terceirapontebot
