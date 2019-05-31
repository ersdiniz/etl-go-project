FROM golang:1.11.5

COPY . /go/src/etl-go-project

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

RUN go get github.com/lib/pq
RUN go get github.com/Nhanderu/brdoc

ADD . /go/src/app

RUN go get -v