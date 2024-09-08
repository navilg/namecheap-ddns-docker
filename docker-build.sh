#!/ussr/bin/env bash

if [ $# -ne 1 ]; then
    echo "Exactly one argument required"
    echo "bash docker-build.sh VERSION"
    echo "  e.g. bash docker-build.sh 1.3.0"
    exit 1
fi

docker buildx create --use --name mybuild
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t linuxshots/namecheap-ddns:$1 --push --pull .
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t linuxshots/namecheap-ddns:latest --push --pull .