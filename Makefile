include .env
export

proto:
	protoc services/api_gateway/pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/auth_service/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/article_and_post/pkg/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/user_service/service/pb/*.proto --go_out=. --go-grpc_out=.
	protoc services/blogsandposts_service/blog_service/pb/*.proto --go_out=. --go-grpc_out=.



sql-gen:
	echo "Enter the file's name or description (Node keep it short):"
	@read INPUT_VALUE; \
	migrate create -ext sql -dir psql/migration -seq $$INPUT_VALUE


# TODO: Make the following changes in db connectons
# 2. SSL mode enable
migrate-up:
	migrate -path psql/migration -database "postgresql://${PSQLUSER}:${PSQLPASS}@${PSQLHOST}:${PSQLPORT}/${PSQLDB}?sslmode=disable" -verbose up

migrate-down:
	migrate -path psql/migration -database "postgresql://${PSQLUSER}:${PSQLPASS}@${PSQLHOST}:${PSQLPORT}/${PSQLDB}?sslmode=disable" -verbose down 1

migrate-force:
	echo "Enter a version:"
	@read INPUT_VALUE; \
	migrate -path psql/migration -database "postgresql://${PSQLUSER}:${PSQLPASS}@${PSQLHOST}:${PSQLPORT}/${PSQLDB}?sslmode=disable" -verbose force $$INPUT_VALUE