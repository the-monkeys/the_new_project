#!/bin/bash

set -x

MONKEY_ROOT="/the_monkeys"

for dir in services/*/cmd
do 
    # Split dir to get the service name
    IFS='/'
    read -ra ADDR <<</$dir
    microservice_name=${ADDR[2]}

    # Merge the dir again to change dir to the cmd
    IFS=' '
    read -ra ADDR <<<$microservice_name

    echo "Build the $microservice_name"
    (cd "$dir" && go build -o "${MONKEY_ROOT}/bin/$microservice_name"); 
done
