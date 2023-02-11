#!/bin/bash

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