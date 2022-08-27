#!/ussr/bin/env bash

docker buildx create --use
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t linuxshots/namecheap-ddns:1.1.0 --push --pull .
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t linuxshots/namecheap-ddns:latest --push --pull .