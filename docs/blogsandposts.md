## Create a blog

```
curl --location --request POST 'https://test-apis.camdvr.org:8080/api/v1/post/create' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU3NjI2MjAsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6NCwiRW1haWwiOiJkYXZldHdlZXRsaXZlQGdtYWlsLmNvbSJ9.K1s0WQ7oyTkD77ZAIDgGAMyWypOzCZxf1CxU0-Kkj6o' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Online HTML Editor    ",
    "content": "<!doctype html>\n<html>\n<head>\n\t<title>HTML Editor - Full Version</title>\n</head>\n<body>\n<h1>HTML Editor - Full Version</h1>\n\n<p>This HTML editor has a larger toolbar than the <a href=\"https://www.quackit.com/html/online-html-editor/\">standard one</a>. And you know what that means? Exactly... more buttons!</p>\n\n<p>Here is some <strong>important text</strong>. And here is text that has been <em>emphasized</em>.</p>\n\n<p>Below is an unordered (bullet) list, and an ordered/numbered list:</p>\n\n<p>Grocery list:</p>\n\n<ul>\n\t<li>Apples</li>\n\t<li>Oranges</li>\n\t<li>Bananas</li>\n</ul>\n\n<p>Things to do:</p>\n\n<ol>\n\t<li>Buy groceries</li>\n\t<li>Get haircut</li>\n\t<li>Write some HTML!</li>\n</ol>\n\n<p>View the source code by clicking on the &quot;Source&quot; button above.</p>\n</body>\n</html>\n",
    "author": "Dave Augustus",
    "author_id": "1",
    "published": true,
    "tags": [
        "HTML",
        "FirstPost ",
        "Wow"
    ]
}'
```