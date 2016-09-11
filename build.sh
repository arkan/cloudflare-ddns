#!/bin/bash

set -e

cd $(dirname $0)

docker build -t=arkan/cloudflare-ddns:latest .
docker push arkan/cloudflare-ddns:latest
