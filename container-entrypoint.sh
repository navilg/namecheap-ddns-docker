#!/usr/bin/env bash

set -e

if [ "$NC_HOST" == "" -o "$NC_DOMAIN" == "" -o "$NC_PASS" == "" ]; then
    echo "ERROR NC_HOST, NC_DOMAIN and NC_PASS are mandatory."
    echo "Use --env with docker run to pass these environment variables."
    exit 1
fi

/app/ncddns --host="$NC_HOST" --domain="$NC_DOMAIN" --password="$NC_PASS"