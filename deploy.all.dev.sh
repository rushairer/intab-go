#!/bin/bash

docker cp intab-apiserver:/go/src gopath/
docker cp intab-webserver:/go/src gopath/
rm -rf gopath/intab*
docker build -f dockerfiles/Dockerfile.intab.all.local -t intab-all .
