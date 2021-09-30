FROM golang:1.15-alpine3.12 AS builder

RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/https\:\/\/alpine.global.ssl.fastly.net/g' /etc/apk/repositories && \
    apk add --no-cache \
      ca-certificates \
      bash \
      gcc \
      libc-dev \
      git

WORKDIR /app
COPY . .

ENV GOPRIVATE=scm.applatform.io

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o api cmd/*.go

FROM alpine:3.12
ENV USER=appuser
ENV UID=10001
EXPOSE 3004

RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/https\:\/\/alpine.global.ssl.fastly.net/g' /etc/apk/repositories \
    && apk --update --no-cache add bash

RUN adduser \
    --disabled-password \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

USER "$USER":"$USER"

COPY scripts /app/scripts
COPY --from=builder /app/api /app/api

ENTRYPOINT ["/app/scripts/entrypoint.sh"]
