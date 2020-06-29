```json
input
{
  "pin": {
    "id": 2,
    "userId": 1,
    "title": "pin1",
    "description": "pin1 description",
    "imageUrl": "https://pinko-bucket.s3-ap-northeast-1.amazonaws.com/pins/009174cf-c4fe-4a49-a4b2-cb24a98887c9.png",
    "isPrivate": false,
    "createdAt": "2020-06-27 10:37:59",
    "updatedAt": "2020-06-27 10:37:59"
  },
  "tags": [
    "neko",
    "inukkoro",
    "sushi"
  ]
}
```
これを受け取って、tagとpinをMySQLで紐付けます。


```json
output
{
  "pin": {
    "id": 2,
    "userId": 1,
    "title": "pin1",
    "description": "pin1 description",
    "imageUrl": "https://pinko-bucket.s3-ap-northeast-1.amazonaws.com/pins/009174cf-c4fe-4a49-a4b2-cb24a98887c9.png",
    "isPrivate": false,
    "createdAt": "2020-06-27 10:37:59",
    "updatedAt": "2020-06-27 10:37:59"
  },
  "tags": [
    {
      "id": 4,
      "tag": "neko"
    },
    {
      "id": 5,
      "tag": "inukkoro"
    },
    {
      "id": 6,
      "tag": "sushi"
    }
  ]
}
```
こういうものを、insertDynamoのlambdaに渡します。

## Usage
execute `./scripts/zip.sh`

upload handler.zip to lambda
