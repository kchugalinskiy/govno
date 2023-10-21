FROM golang:1.21-alpine AS builder

COPY . /go/src/github.com/kchugalinskiy/education

RUN apk update && apk add --no-cache git && \
    go build -o /go/bin/echo-server /go/src/github.com/kchugalinskiy/education/main.go

FROM alpine:latest

COPY --from=builder /go/bin/echo-server /usr/local/bin/echo-server

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/echo-server"]