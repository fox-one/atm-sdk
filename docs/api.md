# Fox ATM Api

* api host: https://efox.fox.one
* request content type is always **application/json**
* atm receipt wallet id: fb151d94-cc0e-358b-9cdb-57215a557a31

## Authorization

Sign a JWT token with the private key

**token payload**

```json5
{
  "uid": "8017d200-7870-4b82-b53f-74bae1d2dad7", // uuid.NewV4().String()
  "exp": 1594637276 // time.Now().Add(exp).Unix
}
```

## Payment

Transfer paying assets to atm receipt wallet id with order memo below.
Snapshot TraceID is the order id.

**Order memo**

```json5
{
  "a": "btc",
  "s": "Market"
}
```

**ATM transfer memo**

```json5
{
  "s": "ATM",
  "c": "FILLED or REFUND",
  "t": "8017d200-7870-4b82-b53f-74bae1d2dad7" // order id
}
```

## Markets

### List Pairs

List all exchange pairs

```http request
GET /api/pairs
```

**Response:**

```json5
{
  "data": [
    {
        "base": {
            "asset_id": "f5ef6b5d-cc5a-3d90-b2c0-a2fd386e7a3c",
            "logo": "https://images.mixin.one/ml7tg1ZGrQt6IJSvEusWFfthosOB98GWN7r4EkmgSP8tbJHxK7OWki9zfZFFDCDOJE0nlLBR6dc4nbUguXE3Bg4=s128",
            "max": "0",
            "min": "0",
            "name": "BOX Token",
            "precision": 8,
            "symbol": "BOX"
        },
        "base_precision": 8,
        "change": "0.0181",
        "display_symbol": "BOX/XIN",
        "exchange": "box",
        "price": "0.0083",
        "price_buy": "0.0084",
        "price_precision": 4,
        "price_sell": "0.0083",
        "quote": {
            "asset_id": "c94ac88f-4671-3976-b60a-09064f1811e8",
            "logo": "https://images.mixin.one/UasWtBZO0TZyLTLCFQjvE_UYekjC7eHCuT_9_52ZpzmCC-X-NPioVegng7Hfx0XmIUavZgz5UL-HIgPCBECc-Ws=s128",
            "max": "4.0852",
            "min": "0",
            "name": "Mixin",
            "precision": 4,
            "symbol": "XIN"
        },
        "quote_precision": 4,
        "strategies": [
            "MARKET"
        ],
        "symbol": "BOXXIN"
    }
  ]
}
```

### Read Pair Detail

Read Pair by symbol

```http request
GET /api/pairs/{symbol}
```

**Response:**

```json5
{
    "data": {
        "base": {
            "asset_id": "c6d0c728-2624-429b-8e0d-d9d19b6592fa",
            "logo": "https://images.mixin.one/HvYGJsV5TGeZ-X9Ek3FEQohQZ3fE9LBEBGcOcn4c4BNHovP4fW4YB97Dg5LcXoQ1hUjMEgjbl1DPlKg1TW7kK6XP=s128",
            "max": "0.000545",
            "min": "0.001",
            "name": "Bitcoin",
            "precision": 6,
            "symbol": "BTC"
        },
        "base_precision": 6,
        "change": "0.00541806",
        "display_symbol": "BTC/USDT",
        "exchange": "bigone",
        "price": "9244.47",
        "price_buy": "9244.79",
        "price_precision": 2,
        "price_sell": "9245.26",
        "quote": {
            "asset_id": "4d8c508b-91c5-375b-92b0-ee702ed2dac5",
            "logo": "https://www.fox.one/assets/coins/USDT(ERC20).png",
            "max": "158.15",
            "min": "0.1",
            "name": "Tether USD",
            "precision": 2,
            "symbol": "USDT"
        },
        "quote_precision": 2,
        "strategies": [
            "MARKET",
            "LIMIT",
            "FOLLOW"
        ],
        "symbol": "BTCUSDT"
    }
}
```

### Read Pair Depth

Read pair depth by symbol

```http request
GET /api/pairs/{symbol}/depth
```

**Response:**

```json5
{
    "data": {
        "asks": [
            {
                "amount": "492.20000000",
                "price": "1.45"
            }
        ],
        "bids": [
            {
                "amount": "492.20000000",
                "price": "1.44"
            }
        ],
        "updated_at": "2020-07-15T05:34:27.253506470Z",
        "version": "1594791267"
    }
}
```

## User

User Authorization required

### Me

Read profile

```http request
GET /api/me
```

**Response:**

```json5
{
    "data": {
        "id": "f49c073a-f92c-4268-ab81-25805d8a1999",
        "name": "player",
        "role": "",
        "broker_id": ""
    }
}
```

### Read Order

Read order by order id

```http request
GET /api/orders/{order_id}
```

**Response:**

```json5
{
    "data": {
        "average_price": "9148.54948843", // average fill price
        "created_at": "2020-06-30T06:44:30Z",
        "discount": "0", // ignore
        "extra_filled_amount": "0",
        "fee_amount": "0.00007752", // fee amount
        "fill_symbol": "BTC",
        "filled_amount": "0.07752265", // filled amount
        "filled_funds": "709.2198", // filled pay amount
        "funds": "709.2198", // pay amount
        "id": "6104cb1c-1a2e-42ae-b6ad-e0fc945e039d", // trace id
        "merchant_id": "f49c073a-f92c-4268-ab81-25805d8a108f",
        "pay_symbol": "USDT",
        "price": "0", // ignore
        "side": "Bid", // Bid or Ask
        "state": "Filled",
        "symbol": "BTCUSDT",
        "updated_at": "2020-06-30T06:44:35Z",
        "user_id": "8017d200-7870-4b82-b53f-74bae1d2dad7"
    }
}
```

### Cancel Order

Cancel Order by order id

```http request
DELETE /api/orders/{order_id}
```

**Response:**

```json5
{
  "data": {}
}
```

### Order History

Read Order history with user id

```http request
GET /api/orders
```

**Parameters:**

```toml
symbol = "BTCUSDT" # optional
side = "ASK or BID" # optional
state = "pending or done"
cursor = "xxx"
limit = 100 # default is 50
```

**Response:**

```json5
{
    "data": {
        "orders": [
            {
                "average_price": "9148.54948843",
                "created_at": "2020-06-30T06:44:30Z",
                "discount": "0",
                "extra_filled_amount": "0",
                "fee_amount": "0.00007752",
                "fill_symbol": "BTC",
                "filled_amount": "0.07752265",
                "filled_funds": "709.2198",
                "funds": "709.2198",
                "id": "6104cb1c-1a2e-42ae-b6ad-e0fc945e039d",
                "merchant_id": "f49c073a-f92c-4268-ab81-25805d8a108f",
                "pay_symbol": "USDT",
                "price": "0",
                "side": "Bid",
                "state": "Filled",
                "symbol": "BTCUSDT",
                "updated_at": "2020-06-30T06:44:35Z",
                "user_id": "8017d200-7870-4b82-b53f-74bae1d2dad7"
            }
        ],
        "pagination": {
            "has_next": true,
            "next_cursor": "112972"
        }
    }
}
```

## Merchant

Merchant Authorization required

### Read Order

Read order by order id

```http request
GET /api/m/orders/{order_id}
```

**Response:**

```json5
{
    "data": {
        "average_price": "9148.54948843", // average fill price
        "created_at": "2020-06-30T06:44:30Z",
        "discount": "0", // ignore
        "extra_filled_amount": "0",
        "fee_amount": "0.00007752", // fee amount
        "fill_symbol": "BTC",
        "filled_amount": "0.07752265", // filled amount
        "filled_funds": "709.2198", // filled pay amount
        "funds": "709.2198", // pay amount
        "id": "6104cb1c-1a2e-42ae-b6ad-e0fc945e039d", // trace id
        "merchant_id": "f49c073a-f92c-4268-ab81-25805d8a108f",
        "pay_symbol": "USDT",
        "price": "0", // ignore
        "side": "Bid", // Bid or Ask
        "state": "Filled",
        "symbol": "BTCUSDT",
        "updated_at": "2020-06-30T06:44:35Z",
        "user_id": "8017d200-7870-4b82-b53f-74bae1d2dad7"
    }
}
```

### Cancel Order

Cancel Order by order id

```http request
DELETE /api/m/orders/{order_id}
```

**Response:**

```json5
{
  "data": {}
}
```

### Order History

Read Order history with user id

```http request
GET /api/m/orders
```

**Parameters:**

```toml
user_id = "a47feefd-e675-3aa1-9c08-a0604d63f8c3" # required
symbol = "BTCUSDT" # optional
side = "ASK or BID" # optional
state = "pending or done"
cursor = "xxx"
limit = 100 # default is 50
```

**Response:**

```json5
{
    "data": {
        "orders": [
            {
                "average_price": "9148.54948843",
                "created_at": "2020-06-30T06:44:30Z",
                "discount": "0",
                "extra_filled_amount": "0",
                "fee_amount": "0.00007752",
                "fill_symbol": "BTC",
                "filled_amount": "0.07752265",
                "filled_funds": "709.2198",
                "funds": "709.2198",
                "id": "6104cb1c-1a2e-42ae-b6ad-e0fc945e039d",
                "merchant_id": "f49c073a-f92c-4268-ab81-25805d8a108f",
                "pay_symbol": "USDT",
                "price": "0",
                "side": "Bid",
                "state": "Filled",
                "symbol": "BTCUSDT",
                "updated_at": "2020-06-30T06:44:35Z",
                "user_id": "8017d200-7870-4b82-b53f-74bae1d2dad7"
            }
        ],
        "pagination": {
            "has_next": true,
            "next_cursor": "112972"
        }
    }
}
```
