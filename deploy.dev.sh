#!/bin/bash
docker build -f dockerfiles/Dockerfile.intab-apiserver.dev -t intab-apiserver-dev .
docker build -f dockerfiles/Dockerfile.intab-webserver.dev -t intab-webserver-dev .
docker build -f dockerfiles/Dockerfile.intab-websocketserver.dev -t intab-websocketserver-dev .
docker build -f dockerfiles/Dockerfile.intab-rpcnodeserver.dev -t intab-rpcnodeserver-dev .
