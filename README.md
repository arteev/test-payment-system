# test-payment-system

* [Swagger](docs/swagger.json)
* [Dependencies](docs/dependencies.md)

## Build and Run

```shell script
docker network create payment_system_network
make build run
```

## Test (unit, integration)

```shell script
    docker network create payment_system_network
    # let's run the dependencies. do not start the service
    make run-local
    make test
```

## Stop

```shell script
    make stop
    docker network rm payment_system_network
```

## Example requests

### New wallet

```shell script
  curl -XPOST "http://localhost:8000/api/v1/payment/wallet" -H "accept: */*" -H "Content-Type: application/json" -d '{"name":"main wallet"}'
  curl -XPOST "http://localhost:8000/api/v1/payment/wallet" -H "accept: */*" -H "Content-Type: application/json" -d '{"name":"second wallet"}'
```
                                     
```json
{
  "data": {
    "id": 1,
    "Name": "main wallet",
    "created_at": 1606037218,
    "updated_at": 1606037218,
    "balance": 0
  },
  "request_id": "0234e1c9-e768-438e-859b-ede6f51e9f24",
  "success": true
}
```
```json
{
  "data": {
    "id": 2,
    "Name": "second wallet",
    "created_at": 1606037411,
    "updated_at": 1606037411,
    "balance": 0
  },
  "request_id": "5c3bf109-f0b3-4746-8287-49f59315043d",
  "success": true
}
```

### Deposit
```shell script
  curl -X POST "http://localhost:8000/api/v1/payment/deposit" -H "accept: */*" -H  "Content-Type: application/json" -d '{"wallet_id": 1, "amount":5.99}'
```

```json
{
  "data": {
    "id": 1,
    "created_at": 1606037619,
    "wallet_id": 1,
    "amount": 5.99
  },
  "request_id": "4e180321-f210-4c49-9cc5-757399cdd797",
  "success": true
}
```

### Get wallet
```shell script
  curl -X GET "http://localhost:8000/api/v1/payment/wallet?id=1"   -H  "accept: */*" -H  "Content-Type: application/json" 
```
```json
{
  "data": {
    "id": 1,
    "Name": "main wallet",
    "created_at": 1606037577,
    "updated_at": 1606037619,
    "balance": 5.99
  },
  "request_id": "c71e2e14-c39b-4fdf-bf54-1090285e3cf2",
  "success": true
}
```
                                                                                                                                                      

### Transfer money
```shell script
  curl -X POST "http://localhost:8000/api/v1/payment/transfer" -H  "accept: */*" -H  "Content-Type: application/json" -d '{"wallet_from": 1,"wallet_to":2, "amount":1.59}'
```

[Check](#Get wallet) the balance of both wallets. The balance of the wallets must change by the transfer amount.

```json
{
  "data": {
    "id": 1,
    "amount": 1.59,
    "created_at": 1606037876
  },
  "request_id": "6d87781f-f34d-4f29-affc-99542482c455",
  "success": true
}
```

### Report "operations". Format: text/csv

Request parameters

|Name|Required|type|
|---|---|---|
| oper |  optional | string:enum(deposit,withdraw) |
| date_from | optional | unix time  | 
| date_to | optional | unix time  | 


```shell script
  curl -X GET "http://localhost:8000/api/v1/payment/operations?oper=deposit&wallet_id=2"   -H  "accept: */*" -H  "Content-Type: text/csv"
``` 

```
    "Date","ID","Wallet ID","Wallet Name","Unit","Operation Type","Amount","Wallet Destination ID","Wallet Destination"
    2020-11-22 09:37:56,3,2,second wallet,transfer,deposit,1.590000,,
```

```shell script
  curl -X GET "http://localhost:8000/api/v1/payment/operations?oper=&wallet_id=1"   -H  "accept: */*" -H  "Content-Type: text/csv" 
```

```
    "Date","ID","Wallet ID","Wallet Name","Unit","Operation Type","Amount","Wallet Destination ID","Wallet Destination"
    2020-11-22 09:33:39,1,1,main wallet,deposit,deposit,5.990000,,
    2020-11-22 09:37:56,2,1,main wallet,transfer,withdraw,1.590000,2,second wallet
```

### Get version 

```shell script
  curl -XGET "http://localhost:8000/api/v1/internal/payment/version"
```

```json
{
  "data": {
    "version": "90f34f1ba213fb237e8505337a31db93bbd131e4",
    "date_build": "2020-11-22_09:21:51",
    "git_hash": "90f34f1ba213fb237e8505337a31db93bbd131e4"
  },
  "request_id": "1f60034f-9042-418a-8e8c-25348dda5a22",
  "success": true
}
```
