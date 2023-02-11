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
IMAGE_NAME=postgres

# Set the Postgres password
POSTGRES_USER=postgres
POSTGRES_PASSWORD=Secret
POSTGRES_DB=the_monkeys

# Create and run the container
docker run -d --name $CONTAINER_NAME \
    -e POSTGRES_USER=$POSTGRES_USER
    -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
    -e POSTGRES_DB=$POSTGRES_DB \
    -p 5432:5432 \
    $IMAGE_NAME


echo "Docker container has been created and running!"

MIGRATION_DIR=psql/migration

sql_files=$(ls $MIGRATION_DIR/*.up.sql)

echo "The following files are set to migrate."
echo $sql_files

# Loop through each SQL file and migrate it to the database
for file in $sql_files
do
  echo "Migrating $file..."

  # Use the docker exec command to run psql in the container and execute the SQL file
  docker exec -i $CONTAINER_NAME psql -U $POSTGRES_USER -d $POSTGRES_DB -v ON_ERROR_STOP=1 -f $file

  # Check the exit code of psql and exit the script if there was an error
  if [ $? -ne 0 ]
  then
    echo "Error migrating $file"
    exit 1
  fi
done

echo "All migrations completed successfully."



