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
POSTGRES_PASSWORD=Secret
POSTGRES_DB=the_monkeys

# Create and run the container
docker run -d --name $CONTAINER_NAME \
  -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  -e POSTGRES_DB=$POSTGRES_DB \
  -p 5432:5432 \
  $IMAGE_NAME