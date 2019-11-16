#!/usr/bin/env bash

#GOOS=linux make build

env GOOS=linux go build -o env2stdout main.go

docker run -it --rm -v `pwd`:/tmp  \
	--env-file .env \
	--network localstack_default \
	gd/awssm2env:latest \
	/tmp/env2stdout
