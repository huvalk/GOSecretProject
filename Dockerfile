FROM golang:alpine

COPY . /app

WORKDIR /app

RUN apk add make && make build-app

CMD sleep 10 && /app/app