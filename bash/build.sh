#!/bin/bash

set -x

for dir in services/*/cmd; 

do 
    # Split dir to get the service name
    IFS='/'
    read -ra ADDR <<</$dir
    microservice_name=${ADDR[2]}

    # Merge the dir again to change dir to the cmd
    IFS=' '
    read -ra ADDR <<<$microservice_name

    echo "Build the $microservice_name"
    (cd "$dir" && go build -o "/usr/local/bin/the_monkeys/$microservice_name"); 
done



# restart services to load the new code changes
sudo systemctl restart microservice_name 