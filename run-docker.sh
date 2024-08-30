#!/usr/bin/env bash

docker build -t foobar -f Dockerfile --platform linux/arm64 .
docker run --rm -ti --platform linux/arm64 foobar /run.sh
