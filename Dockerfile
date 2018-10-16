FROM golang:latest

RUN mkdir /go/src/app

COPY . /go/src/app

WORKDIR /go/src/app/pb

RUN go get -v ./... \
    && cd .. \
    && go get -d -v ./...

CMD cd ../ \
    && go run server.go