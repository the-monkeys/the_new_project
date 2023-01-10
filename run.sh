#!/bin/bash
go run api_gateway/cmd/main.go 
go run article_and_post/cmd/main.go
go run auth_service/cmd/main.go

