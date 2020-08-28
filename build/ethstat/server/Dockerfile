FROM golang:1.15-alpine as builder
RUN apk add --update git alpine-sdk

ENV GO111MODULE=on

WORKDIR /app

COPY ../../.. .

RUN go get -u github.com/gobuffalo/packr/packr && go mod download && packr

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine
RUN apk add --update --no-cache ca-certificates openssl postgresql-client
COPY --from=builder /app/ethstat /app/
COPY --from=builder /app/build/ethstat/server/docker-entrypoint.sh /app/

RUN chmod +x /app/docker-entrypoint.sh && \
    chmod +x /app/ethstat

EXPOSE 8080

ENTRYPOINT ["/app/docker-entrypoint.sh"]

CMD ["/app/ethstat", "server"]