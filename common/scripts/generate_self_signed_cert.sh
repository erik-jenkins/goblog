#!/bin/bash

# set working directory to common directory
cd "$(dirname "$0")"/..

# create tls directory if it doesn't exist
if [[ ! -d "tls" ]]; then
  mkdir tls
fi

# change to tls directory
cd tls

# generate self signed cert
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

# go back to common directory
cd ..

# copy to api and ui folders
cp -r tls ../api
cp -r tls ../ui
