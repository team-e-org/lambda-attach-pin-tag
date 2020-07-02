```json
input
{
  "pin": {
    "id": 2,
    "userId": 1,
    "title": "pin1",
    "description": "pin1 description",
    "imageUrl": "https://d1khmj8lxb1rtx.cloudfront.net/pins/1/4184f0b9-a5c4-4343-b6d7-14c75d1775ae.jpg",
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
    "imageUrl": "pins/1/4184f0b9-a5c4-4343-b6d7-14c75d1775ae.jpg",
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

`aws lambda update-function-code --function-name attachTag --zip-file fileb://handler.zip`
