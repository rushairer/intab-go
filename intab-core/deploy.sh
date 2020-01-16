#!/bin/bash
#docker rm -f intab-app
docker build -t intab:latest .
#docker run -d -it -p 80:8080 --name intab-app intab:latest
