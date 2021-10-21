#!/usr/bin/env bash

# Just an example how to quickly run the server locally with environment variables set.
export MYSQL_DATABASE=1
export MYSQL_USERNAME=fred
export MYSQL_PASSWORD=123
export MYSQL_HOST_URL=example
export MYSQL_HOST_PORT=1
export RDS_SSL_CERT=A
export HTTP_CLIENT_TIMEOUT=89s
export HEADER_TIMEOUT=44s
export CUCKOO_SERVICES="A,B,C,D E F"
export ENABLE_PAYID=1
export FACTORY=prelive
export MAX_CONN_LIFETIME=56m
export ASYNC_REQUEST_QUEUE_URL=123
export WAPI_QUEUE_URL=456

go run .
