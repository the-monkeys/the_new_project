#!/bin/bash

# Delete all the .pb.go files from all the directories.

# Generate code from proto files
# protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.

for files in $(ls -d */)
do
    protoc $files pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
done