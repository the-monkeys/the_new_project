# Contributing Guidelines
We're glad you're thinking about contributing to The Monkeys. If you think something is missing or could be improved, please open issues and pull requests. If you'd like to help this project grow, we'd love to have you. 

* If you find any issue or bug please create a Github issue or mail us at [mail.themonkeys.life@gmail.com](mail.themonkeys.life@gmail.com). 
* Create branches in your fork, and submit PRs from your forked branch.

# Local Setup Requirement
* Postgres
* Opensearch
* Golang 1.18
* Protoc compiler

NOTE: Have a config file in /etc/the_monkeys/dev.env
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



### NOTE:
* Raise a pull request(PR) against the main branch, rebase before raising the PR.
* All the PRs need to be approved by [Dave Augustus](https://github.com/daveaugustus) before the merge.
* Code consistency needs to be checked before raising the PR.
* Spelling needs to be checked before the PR.
* The sensitive information like environment variables shouldn't be in the code.
* Linting needs to be checked.

