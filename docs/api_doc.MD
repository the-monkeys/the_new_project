## Register API:
```
curl --location --request POST 'http://localhost:5001/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Irak",
    "last_name": "Rigia",
    "email": "vir@beta.com",
    "password": "1234"
}'

```

## Login API:
```
curl --location --request POST 'http://localhost:5001/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "vir@beta.com",
    "password": "1234"
}'
```

## Get blogs 
```
curl --location --request GET 'http://localhost:5001/api/v1/post/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "alpha@beta.com",
    "password": "1234"
}'

```
## Get blogs by Id
```
curl --location --request GET 'http://localhost:5001/api/v1/post/121778fb-dfe5-4671-926a-85809762920f' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "alpha@beta.com",
    "password": "1234"
}'
```

## Create an article
```
curl --location --request POST 'http://localhost:5001/api/v1/post/create' \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDYyOTE4NTAsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MCwiRW1haWwiOiJ2aXJAYmV0YS5jb20ifQ.dO6mzGOkZ8uimThIOgdolJPGFXXhu75oOypmxRclUaM' \
--header 'Content-Type: application/json' \
--data-raw '{
        "title": "Article 3    ",
        "content": "<!doctype html>\n<html>\n<head>\n\t<title> Article 3 - Full Version</title>\n</head>\n<body>\n<h1>HTML Editor - Full Version</h1>\n\n<p>This HTML editor has a larger toolbar than the <a href=\"https://www.quackit.com/html/online-html-editor/\">standard one</a>. And you know what that means? Exactly... more buttons!</p>\n\n<p>Here is some <strong>important text</strong>. And here is text that has been <em>emphasized</em>.</p>\n\n<p>Below is an unordered (bullet) list, and an ordered/numbered list:</p>\n\n<p>Grocery list:</p>\n\n<ul>\n\t<li>Apples</li>\n\t<li>Oranges</li>\n\t<li>Bananas</li>\n</ul>\n\n<p>Things to do:</p>\n\n<ol>\n\t<li>Buy groceries</li>\n\t<li>Get haircut</li>\n\t<li>Write some HTML!</li>\n</ol>\n\n<p>View the source code by clicking on the &quot;Source&quot; button above.</p>\n</body>\n</html>\n",
        "author": "Alpha Beta",
        "author_id": "5",
        "published": true,
        "tags": [
            "HTML",
            "FirstPost ",
            "Wow"
        ]
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

## Edit Article 
```
curl --location --request PATCH 'http://localhost:5001/api/v1/article/post/a0b81286-bfdd-49f1-8bb3-56ace2918736' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ5MDAzMzksImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MywiRW1haWwiOiJraGFzYmFAbWFpbC5jb20ifQ.p8oKo8Ny1cSh67p_C7_oagTNdnIUpQxosz1sUSG4--k' \
--data-raw '{
    "title": "Changed title 11"

}'
```