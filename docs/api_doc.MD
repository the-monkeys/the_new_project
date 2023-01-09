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