#!/bin/bash
#docker rm -f intab-server
docker build -f Dockerfile.dev -t intab-dev .
#docker run -d -it -p 80:8080 -v "$PWD":/go/src/intab --name intab-server intab-dev
docker build -f Dockerfile.websocket.dev -t intab-websocket-dev .
