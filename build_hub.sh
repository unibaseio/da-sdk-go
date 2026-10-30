#!/bin/sh
# Build and push the hub image. IMAGE defaults to the repository the hub
# compose file pulls (docker-compose/docker-compose-hub.yml).
set -e
version=${1:?usage: build_hub.sh <version>}
IMAGE=${IMAGE:-unibaseio/hub}

docker build -t "$IMAGE:hub-$version" -f Dockerfile .
docker tag "$IMAGE:hub-$version" "$IMAGE:hub-latest"
docker push "$IMAGE:hub-$version"
docker push "$IMAGE:hub-latest"
