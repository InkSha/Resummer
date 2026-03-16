#!/bin/sh

cd server

swag init  --parseInternal -g ./cmd/server/main.go -o ./cmd/server/docs/

go run cmd/server/main.go
