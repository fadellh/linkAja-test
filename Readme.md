# LinkAja Test

## Run Project with Docker

```sh
docker-compose up --build
```

## Run the test

```sh
go test -v -cover ./...
```

# REST API

## Check Saldo

Untuk mendapatkan saldo user saat ini

### Request

`GET /account/{account_number}`

    curl --location --request GET 'http://localhost:2801/account/555002'

## Transfer

Untuk melakukan transfer balance dari akun user satu ke akun user lainnya

### Request

`POST /account/{from_account_number}/transfer`

```
curl --location --request POST 'http://localhost:2801/account/555001/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "to_account_number": "555002",
    "amount": 100
}'
```
