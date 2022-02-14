FROM golang:alpine

ADD main.go /go/src/hello-app/main.go

RUN cd /go/src/hello-app && \
    go mod init && \
    go install /go/src/hello-app

FROM alpine:latest

ADD hello.html /www
ADD main.css /www

COPY --from=0 /go/bin/hello-app .
ENV PORT 8080
CMD ["./hello-app"]