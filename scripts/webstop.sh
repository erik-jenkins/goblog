#!/bin/bash

# set working directory to project root
cd "$(dirname "$0")"/..

docker-compose -f docker-compose.dev.yml down
[ -f docker-compose.debug.yml ] && docker-compose -f docker-compose.dev.yml -f docker-compose.debug.yml down
