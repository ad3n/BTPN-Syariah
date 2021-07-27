FROM golang:alpine

RUN apk add musl-dev libc-dev gcc
RUN mkdir -p /go/src/resto
ADD . /go/src/resto
WORKDIR /go/src/resto

RUN go get
RUN go build -o app .

EXPOSE 3000

CMD ["/go/src/resto/app"]
