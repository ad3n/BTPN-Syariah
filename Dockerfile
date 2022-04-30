FROM golang:alpine as builder

RUN apk add musl-dev libc-dev gcc git
RUN mkdir -p /go/src/resto
ADD . /go/src/resto
WORKDIR /go/src/resto

RUN go get
RUN go test -coverprofile /tmp/coverage ./... -v
RUN go build -o app .

FROM alpine:latest
COPY --from=builder /go/src/resto/app /usr/local/bin/resto

EXPOSE 3000

CMD ["/usr/local/bin/resto"]
