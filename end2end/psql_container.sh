#!/bin/bash

set -x

# Get a list of all container IDs
container_ids=$(docker ps -aq)

# Remove all containers
if [ ! -z "$container_ids" ]
then
  docker rm -f $container_ids
  echo "Removed all containers."
else
  echo "No containers to remove."
fi

# Define the container name and image name
CONTAINER_NAME=subtle_art
IMAGE_NAME=postgres:12-alpine

OPENSEARCH_CONTAINER=art_of_writers
OPENSEARCH_IMAGE=opensearchproject/opensearch:latest
# Set the Postgres password
POSTGRES_USER=root
POSTGRES_PASSWORD=Secret
POSTGRES_DB=the_monkeys

# Create and run the container
docker run -d --name $CONTAINER_NAME \
    -e POSTGRES_USER=$POSTGRES_USER \
    -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
    -e POSTGRES_DB=$POSTGRES_DB \
    -p 5432:5432 \
    $IMAGE_NAME


sudo docker run -d --name $OPENSEARCH_CONTAINER -p 9200:9200 -p 9600:9600 -e "discovery.type=single-node" $OPENSEARCH_IMAGE


echo "Docker containers have been created and running!"

MIGRATION_DIR=psql/migration

# Install Golang Migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate.linux-amd64 /usr/local/bin/migrate
sudo chmod +x /usr/local/bin/migrate

migrate -path psql/migration -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@127.0.0.1:5432/$POSTGRES_DB?sslmode=disable" -verbose up

# docker exec $CONTAINER_NAME mkdir -p $MIGRATION_DIR
# docker cp psql/migration/. $CONTAINER_NAME:/psql/migration

# # Migrate the SQL files in order
# for FILE in $(ls psql/migration/*.up.sql | sort); do
#     echo "Migrating file $FILE"
# #   docker exec -i $CONTAINER_NAME psql -U $POSTGRES_USER -d $POSTGRES_DB -v ON_ERROR_STOP=1 -f $FILE
#   docker exec  $CONTAINER_NAME psql -U $POSTGRES_USER -d $POSTGRES_DB -p $POSTGRES_PASSWORD -p 5431 -f $FILE

# done


# sql_files=$(ls $MIGRATION_DIR/*.up.sql)

# echo "The following files are set to migrate."
# echo $sql_files

# # Loop through each SQL file and migrate it to the database
# for file in $sql_files
# do
#   echo "Migrating $file..."

#   # Use the docker exec command to run psql in the container and execute the SQL file
#   docker exec -i $CONTAINER_NAME psql -U $POSTGRES_USER -d $POSTGRES_DB -v ON_ERROR_STOP=1 -f $file
#     # docker exec -it $CONTAINER_NAME psql -U $POSTGRES_USER -d $POSTGRES_DB -f $file


#   # Check the exit code of psql and exit the script if there was an error
#   if [ $? -ne 0 ]
#   then
#     echo "Error migrating $file"
#     exit 1
#   fi
# done

echo "All migrations completed successfully."



