# なんかよくわからないけど動いてる

FROM golang:1.16.14-alpine3.15 AS dev

WORKDIR /go/src/app

RUN go get github.com/cespare/reflex
CMD reflex -r '(\.go$|go\.mod)' -s go run cmd/main.go


FROM golang:1.16.14-alpine3.15 AS builder
WORKDIR /go/src/github.com/do-it-if-i-can/server

# go.modとdo.sumだけコピーしてもいいが上手くいかないので全部コピーする
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/main.go


FROM alpine:3.12.3 AS prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/do-it-if-i-can/server/app .
CMD ["/root/app"]