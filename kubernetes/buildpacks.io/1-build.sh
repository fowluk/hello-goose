#!/bin/bash

# Go to the right place
cd $(dirname $0)/../../python

# Build Hello Goose
pack build hello-goose-cnb

# Push the image to Docker Hub - replace the repo with yours!
docker tag hello-goose-cnb jjeffreypivotal/hello-goose-cnb
docker push jjeffreypivotal/hello-goose-cnb
