### Upload Profile details:
```
curl --location --request POST 'https://localhost:5001/api/v1/profile/user' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc0MjU3MzUsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.BYYSo6hPMzTY6M1dB7JCIP0r6UWVSDfbo7gnIaOqBfY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Dave",
    "last_name": "Augustus",
    "country_code": "+91",
    "mobile": "9611808267",
    "about": "Founder of the monkeys inc.",
    "instagram": "_the__monkeys_",
    "twitter": "TheMonkeysLife",
    "email": "mail.themonkeys.life@gmail.com"
}'

```

### Get Profile pics
```
curl --location --request GET 'https://localhost:5001/api/v1/profile/user/2' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc0MjU3MzUsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJtYWlsLnRoZW1vbmtleXMubGlmZUBnbWFpbC5jb20ifQ.BYYSo6hPMzTY6M1dB7JCIP0r6UWVSDfbo7gnIaOqBfY' \
--form 'image=@"/home/dave/Pictures/Screenshots/nginx_ssl.png"'
```