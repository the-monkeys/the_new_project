
### To setup this locally in linux machine, lease install
* Postgres
* Opensearch
* Golang
* Protoc compiler

NODE: Have a config file in /etc/the_monkeys/dev.env
```
PORT=0.0.0.0:5001
AUTH_SVC_URL=localhost:50051
STORY_SVC_URL=localhost:50052
USER_SVC_URL=localhost:50053


AUTH_SERVICE_PORT=:50051
DB_URL=postgres://user:password@localhost:5432/dbname
JWT_SECRET_KEY=secret_key

ARTICLE_SERVICE_PORT=:50052
OPENSEARCH_ADDRESS=https://localhost:9200
OSUSERNAME=username
OSPASSWORD=password

USER_SERVICE_PORT=:50053

```




# Install linting tool
```
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z
```

## Run go lint command
$ `golangci-lint run`