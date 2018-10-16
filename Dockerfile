FROM golang:latest

RUN mkdir /go/src/app

COPY . /go/src/app

WORKDIR /go/src/app/

RUN go get -v ./...

CMD go run server.go