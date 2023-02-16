#!/bin/bash

set -x

echo "Generate server side self-signed certificates"
openssl genrsa -out certs/server.key 2048

openssl req -new -key certs/server.key -out certs/server.csr
openssl x509 -req -days 365 -in certs/server.csr -signkey certs/server.key -out certs/server.crt -outform PEM
echo "Generated server side self-signed certificates"



# TODO: remove this file if not in use