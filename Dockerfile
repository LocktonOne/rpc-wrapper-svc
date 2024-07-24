FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokene/lockton-one/rpc-wrapper-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/rpc-wrapper-svc /go/src/gitlab.com/tokene/lockton-one/rpc-wrapper-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/rpc-wrapper-svc /usr/local/bin/rpc-wrapper-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["rpc-wrapper-svc"]
