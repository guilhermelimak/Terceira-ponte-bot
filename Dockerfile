FROM golang:latest

RUN mkdir /app

COPY . /app

WORKDIR /app

ENTRYPOINT /app/terceirapontebot
