# the_new_project


### To setup this locally in linux machine, lease install
* Postgres
* Opensearch

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


# The current architecture diagram.
![alt text](docs/current_architecture.jpg)


# Contribution Guidelines:
If you are want to contribute to the project, mail us at davetweetlive@gmail.com for access with your github account details. 

* Once you get the access, pick an issue and code for it.
* Raise a pull request(PR) against the main branch, rebase before raising the PR.
* All the PRs need to be approved by [Dave Augustus](https://github.com/daveaugustus) before the merge.
* Code consistency needs to be checked before raising the PR.
* Spelling needs to be checked before the PR.
* The sensitive information like environment variables shouldn't be in the code.