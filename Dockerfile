FROM golang:1.6

USER nobody

RUN mkdir -p /go/src/github.com/agajdosi/golang-ex/webserver
WORKDIR /go/src/github.com/agajdosi/golang-ex/webserver

COPY . /go/src/github.com/agajdosi/golang-ex/webserver
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]
