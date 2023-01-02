#!/bin/bash
go run api_gateway/cmd/main.go 
go run article_and_post/cmd/main.go 

protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.