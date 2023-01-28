proto:
	protoc services/api_gateway/pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/auth_service/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/article_and_post/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/user_profile/user_service/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/blogsandposts_service/blog_service/pb/*.proto --go_out=. --go-grpc_out=.



sql-gen:
	migrate create -ext sql -dir db -seq init_post_schema


all:
    echo ${PATH_SQL_MIGRATE}

# TODO: Make the following changes in db connectons
# 1. Remove the sensitive information to secret service
# 2. SSL mode enable

# Keeping temproraily as it's the local configuration.
migrate-up:
	migrate -path db -database "postgresql://root:Secret@localhost:5432/the_monkeys?sslmode=disable" -verbose up