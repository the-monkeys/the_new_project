## Register API

### Request
```
curl --location --request POST 'https://localhost:8080/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Firstname",
    "last_name": "Lastname",
    "email": "username@gmail.com",
    "password": "1234"
}'
```

### Response
```
{
    "status": 201,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkzMzA2ODYsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJkYXZldHdlZXRsaXZlQGdtYWlsLmNvbSJ9.hDKzD_s-e8kfMRgyX-AIzqyPhKL2Xogsx7r0EYMS2Rw"
}
```
NOTE: The register API take email as the username, once registration id done, it send a verification email which needs to be verified by the user.


## Login API

### Request
```
curl --location --request POST 'https://localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "username@gmail.com",
    "password": "1234"
}'
```

### Response
```
{
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkzMzE1NTksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.C_VkEUhlVyyMm97wU3OQPoiOBlYDRL2sPkN_gkMMtNY"
}
```