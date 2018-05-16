FROM golang as builder
WORKDIR /go/tg-markdown-finder
COPY ./src ./src
ENV GOPATH /go/tg-markdown-finder/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tg-markdown-finder ./src/finder/bot/bot.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /root/
COPY --from=builder /go/tg-markdown-finder/tg-markdown-finder .
COPY docker/entrypoint.sh /tmp/entrypoint.sh
ENTRYPOINT ["/tmp/entrypoint.sh"]