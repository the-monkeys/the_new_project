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


## Password Reset
Password reset has 3 different apis:

    - API which sends forgot password request.
    - API which can be triggered from email verification URL
    - API which takes a new password.

### Forgot password request
Request
```
curl --location --request POST 'https://localhost:8080/api/v1/auth/forgot-pass' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "davetweetlive@gmail.com"
}'
```

Response
```
Get an email containing a verification URL
```

### Verify Email
Request
```
curl --location --request GET 'https://localhost:8080/api/v1/auth/reset-password?verificationparams'
```

Response
```
{
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0Njg3ODMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MjMsIkVtYWlsIjoiZGF2ZXR3ZWV0bGl2ZUBnbWFpbC5jb20ifQ.4yZt3q54EE0fls8JbUZdldQ1QY0e1NIQVELetw7p3wU"
}
```

### Reset Password
Request
```
curl --location --request POST 'https://localhost:8080/api/v1/auth/update-password' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0Njg3ODMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MjMsIkVtYWlsIjoiZGF2ZXR3ZWV0bGl2ZUBnbWFpbC5jb20ifQ.4yZt3q54EE0fls8JbUZdldQ1QY0e1NIQVELetw7p3wU' \
--header 'Content-Type: application/json' \
--data-raw '{
    "password": "abcd"
}'
```

Response
```
{
    "status": 200
}
```


## Email Verification
Request
```
curl --location --request POST 'https://localhost:8080/api/v1/auth/req-email-verification' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0Njg3ODMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MjMsIkVtYWlsIjoiZGF2ZXR3ZWV0bGl2ZUBnbWFpbC5jb20ifQ.4yZt3q54EE0fls8JbUZdldQ1QY0e1NIQVELetw7p3wU' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "mail.themonkeys.life@gmail.com"
}'
```

Response
```
{
    "status": 200,
    "message": "Check your email and click on the verify link"
}
```

## Verify Email from clicking Email link
Request
```
curl --location --request POST 'https://localhost:8080/api/v1/auth/verify-email?user=mail.themonkeys.life@gmail.com&evpw=WodTB2rJ8SobMgQ1nrtR245jxOrsovFiKVY9DYwAq9CuixIQ9sziMSyLhJjL97xH'
```
Response
```
{
    "status": 200
}
```