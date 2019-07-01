FROM alpine:3.9
COPY go-echoserver /go-echoserver
ENTRYPOINT ["/go-echoserver"]
