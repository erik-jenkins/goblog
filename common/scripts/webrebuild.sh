#!/bin/bash

# set working directory to common folder
cd "$(dirname "$0")"/..

docker-compose -f docker-compose.dev.yml build
docker-compose -f docker-compose.dev.yml up -d
