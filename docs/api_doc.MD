## Register API:
```
curl --location --request POST 'http://localhost:5001/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Dave",
    "last_name": "Augustus",
    "email": "dave@the_monkey.com",
    "password": "12345"
}'

```

## Login API:
```
curl --location --request POST 'http://localhost:5001/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "dave@the_monkey.com",
    "password": "12345"
}'

```

## Create an article
```
curl --location --request POST 'http://localhost:5001/api/v1/article/' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ5MDAzMzksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MywiRW1haWwiOiJraGFzYmFAbWFpbC5jb20ifQ.p8oKo8Ny1cSh67p_C7_oagTNdnIUpQxosz1sUSG4--k' \
--header 'Content-Type: application/json' \
--data-raw '{
 "title": "Saving to Pocket via Email",
 "content": "Pocket can only save links from emails. ",
 "author": "Dave Augustus",
 "is_draft": true,
 "tags": ["firstpost", "test", "dev"]
}'

```

## View Profile API
```

curl --location --request GET 'http://localhost:5001/api/v1/profile/user' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ5MDAzMzksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MywiRW1haWwiOiJraGFzYmFAbWFpbC5jb20ifQ.p8oKo8Ny1cSh67p_C7_oagTNdnIUpQxosz1sUSG4--k' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1
}'

```

## Get Dashboard Articles
```
curl --location --request GET 'http://localhost:5001/api/v1/article/'
```

## Get Article by Id
```
curl --location --request GET 'http://localhost:5001/api/v1/article/a0b81286-bfdd-49f1-8bb3-56ace2918736'
```