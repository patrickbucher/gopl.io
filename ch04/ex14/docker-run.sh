#!/bin/sh

if [ -z "$1" ]; then
    echo "Missing Search Parameter"
    exit 1
fi
docker run -it -p 8000:8000 --entrypoint "/go/bin/issueserver" issueserver "$1"
