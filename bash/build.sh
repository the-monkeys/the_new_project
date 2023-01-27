#!/bin/bash

source ${MONKEY_SCRIPTS}/common.sh

set -x

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
<<<<<<<< HEAD:scripts/build.sh
    (cd "$dir" && go build -o "${MONKEY_ROOT}/bin/$microservice_name"); 
done
========
    (cd "$dir" && go build -o "/usr/local/bin/the_monkeys/$microservice_name"); 
done



# restart services to load the new code changes
<<<<<<< HEAD
sudo systemctl restart microservice_name 
>>>>>>>> 9e7d540 (SQL script ready):bash/build.sh
=======
>>>>>>> 08c6600 (Build script)
