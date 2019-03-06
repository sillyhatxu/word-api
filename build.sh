#!/usr/bin/env bash

echo build start
go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
docker build -t word-api .
docker tag word-api:latest xushikuan/word-api:1.0
docker push xushikuan/word-api:1.0
echo build end