FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/go-postgres/
COPY . /go/src/github.com/go-postgres/
RUN go build -o ./dist/go-postgres

FROM alpine:3.5
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

WORKDIR /app
COPY ./config/go-postgres.yaml /var/app/
COPY ./config/go-postgres.yaml /
COPY --from=builder go/src/github.com/go-postgres/dist/go-postgres .
EXPOSE 9090
ENTRYPOINT ["./go-postgres"]
