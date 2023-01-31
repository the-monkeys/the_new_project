# Contributing Guidelines
We're glad you're thinking about contributing to The Monkeys. If you think something is missing or could be improved, please open issues and pull requests. If you'd like to help this project grow, we'd love to have you. 

* If you find any issue or bug please create a Github issue or mail us at [mail.themonkeys.life@gmail.com](mail.themonkeys.life@gmail.com). 
* Create branches in your fork, and submit PRs from your forked branch.

# Local Setup Requirement
* Postgres
* Opensearch
* Golang 1.18
* Protoc compiler

NOTE: Have a config file in `/etc/the_monkey/dev.env` if you are using Linux/Mac. In case if you have a Windows machine you can keep the dev.env in your fav directory and set up the path `services/api_gateway/config/config.go` file and other `config.go` files in different microservice.

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

### Run go lint command
$ `golangci-lint run`



# PR Approval and Merge

We are keeping some checks before merging the PRs to the main branch to maintain the code for a long time and for now we have setup the following rules, it future we may update the rules use some automation and pipeline for code and consistency checks.

* All the PRs need to be approved by [Dave Augustus](https://github.com/daveaugustus) before the merge.
* Code consistency needs to be checked before raising the PR.
* Spelling needs to be checked before the PR.
* The sensitive information like environment variables shouldn't be in the code.
* Linting needs to be checked.

