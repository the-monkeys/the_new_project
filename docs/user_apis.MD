### Upload Profile details:
```
curl --location --request POST 'https://localhost:5001/api/v1/profile/user/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2NTg3NTIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.6tAe7sgLJCkwPsQ9ZZJ_TZaPqVr4E7ZNihjD-0Md1nc' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Dave",
    "last_name": "Augustus",
    "country_code": "+91",
    "mobile": "55665548",
    "about": "The Monkeys inc",
    "instagram": "isndfs",
    "twitter": "adsfddd",
    "email": "daedsdfsd@gmail.com"
}'

```

### Get My Profile details:
```
curl --location --request GET 'https://localhost:5001/api/v1/profile/user/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2NTg3NTIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.6tAe7sgLJCkwPsQ9ZZJ_TZaPqVr4E7ZNihjD-0Md1nc' \
--data-raw ''

```

### Upload profile pic
```
curl --location --request POST 'https://localhost:5001/api/v1/profile/user/pic/2' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2NTg3NTIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.6tAe7sgLJCkwPsQ9ZZJ_TZaPqVr4E7ZNihjD-0Md1nc' \
--form 'image=@"/home/dave/Pictures/signal-2023-01-26-235210_002.jpeg"'
```

### Get Profile pics
curl --location --request GET 'https://localhost:5001/api/v1/profile/user/pic/2' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2NTg3NTIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.6tAe7sgLJCkwPsQ9ZZJ_TZaPqVr4E7ZNihjD-0Md1nc' \
--form 'image=@"/home/dave/Pictures/signal-2023-01-26-235210_002.jpeg"'
```