#!/usr/bin/env bash

set -e

if [ "$NC_HOSTS" == "" -o "$NC_DOMAIN" == "" -o "$NC_PASS" == "" ]; then
    echo "ERROR NC_HOSTS, NC_DOMAIN and GD_PASS are mandatory."
    echo "Use --env with docker run to pass these environment variables."
    exit 1
fi

/app/ncddns --host="$NC_HOSTS" --domain="$NC_DOMAIN" --password="$NC_PASS"