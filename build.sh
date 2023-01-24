#!/bin/bash


for dir in services/*/cmd; 

do 
    # Split dir to get the service name
    IFS='/'
    read -ra ADDR <<</$dir
    microservice_name=${ADDR[2]}

    # Merge the dir again to change dir to the cmd
    IFS=' '
    read -ra ADDR <<<$microservice_name

    (cd "$dir" && go build -o "$microservice_name"); 
done



# restart services to load the new code changes