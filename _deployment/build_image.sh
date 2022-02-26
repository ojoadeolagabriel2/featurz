#!/usr/bin/env sh

set -ex

MY_PATH=$(dirname "$0")
MY_PATH=$(cd "$MY_PATH" && cd .. && pwd)
CONTAINER_NAME=featurz

export PATH=$PATH:/usr/local/go/bin

cd "$MY_PATH"
go env
go mod download
# go test -v "$MY_PATH"/...
go build

docker system prune -af
docker build -t app/featurz:latest .

docker stop "$CONTAINER_NAME" || true && docker rm "$CONTAINER_NAME" || true
docker run --name "$CONTAINER_NAME" -p 12345:12345 -d app/featurz:latest