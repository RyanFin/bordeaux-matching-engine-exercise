# bordeaux-matching-engine-exercise
Bordeaux Tech Test

## Instructions
- Run the application with: `$ make run` in the root directory
- Test the application with: `$ make test` in the root directory


## Improvements
- Use of mutexes to prevent race conditions and control access to a resource with the use of locks.

## API

- View Orderbook [GET]: curl --location 'http://localhost:8080/orderbook'

- Limit Buy [POST]: 
```
curl --location 'http://localhost:8080/order' \
--header 'Content-Type: application/json' \
--data '{
  "type": "limit",
  "side": "buy",
  "price": 98,
  "quantity": 1
}'
```

- Limit Sell [POST]:
```
curl --location 'http://localhost:8080/order' \
--header 'Content-Type: application/json' \
--data '{
  "type": "limit",
  "side": "sell",
  "price": 90.0,
  "quantity": 1
}'
```

- Market Buy [POST]:
```
curl --location 'http://localhost:8080/order' \
--header 'Content-Type: application/json' \
--data '{
  "type": "market",
  "side": "buy",
  "price": 20.0,
  "quantity": 1
}'
```

- Market Sell [POST]:
```
curl --location 'http://localhost:8080/order' \
--header 'Content-Type: application/json' \
--data '{
  "type": "market",
  "side": "sell",
  "price": 10.0,
  "quantity": 4
}'
```

