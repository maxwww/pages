FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY . /github.com/maxwww/pages/
WORKDIR /github.com/maxwww/pages/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/maxwww/pages/.bin/bot .
COPY --from=0 /github.com/maxwww/pages/configs configs/

CMD ["./bot"]