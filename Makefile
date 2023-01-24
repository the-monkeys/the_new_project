proto:
	protoc services/api_gateway/pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/auth_service/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/article_and_post/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/user_profile/user_service/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/blogsandposts_service/blog_service/pb/*.proto --go_out=. --go-grpc_out=.


migrate-up:
	migrate -path db/user/migration -database "postgresql://root:Secret@localhost:5432/the_monkeys?sslmode=disable" -verbose up