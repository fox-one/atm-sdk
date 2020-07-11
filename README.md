# atm-sdk
fox atm sdk

* api host `https://efox.fox.one`

## 商户配置

联系 Fox.ONE 获取商户 id，签名私钥，下单 mixin 钱包地址

## 下单

* 构造下单 memo，带上想买的币的 symbol/asset_id 和所使用的策略
* 转账付款的币给指定下单地址
* 转账的 `TraceID` 即这笔订单的订单 id，转账人的钱包 id 为这笔订单的用户 id

**Memo**

```go
action := OrderAction{
    Asset: "BTC", // symbol or asset id
    Strategy: StrategyMarket, // 默认是 StrategyMarket
}
memo := action.Encode()
```

## 查询订单

下单之后可以开始用下单转账的 `TraceID` 作为订单 id 查询订单。
因为异步到账的原因，可能会在一小段时间内查询不到订单

```go
const host = "https://efox.fox.one"
client := NewMerchantServiceProtobufClient(host, &http.Client{})

token := GenerateToken(merchantID, key, time.Minute)
ctx := WithToken(context.Backgroud(), token)

orderID := "<trace id>"
order, err := client.ReadOrder(ctx, &MerchantServiceReq_ReadOrder{
    TraceId: orderID,
})

if err != nil {
    // 查询订单 404，可能还没到账，等待一会再继续查询
    if IsErrorNotFound(err) {
        // handle order not found error
    }

    // handle err
}
```

## 撤单

查询到订单之后，才可以撤单

```go
token := GenerateToken(merchantID, key, time.Minute)
ctx := WithToken(context.Backgroud(), token)

if _, err := client.CancelOrder(ctx, &MerchantServiceReq_CancelOrder{
    TraceId: orderID,
}); err != nil {
    // handle cancel order failed error
}
```

## 查询用户历史订单

```go
token := GenerateToken(merchantID, key, time.Minute)
ctx := WithToken(context.Backgroud(), token)

results, err := client.ListOrders(ctx, &MerchantServiceReq_ListOrders{
    Symbol:   "BTCUSDT",      // pay symbol，为空查询所有交易对
    Side:     "ASK",          // ASK or BID，为空查询所有方向
    Strategy: StrategyMarket, // 策略，为空查询所有策略
    State:    "pending",      // pending or done，为空查询所有状态
    UserId:   order.UserId,   // 用户 id
    Order:    SortOrder_ASC,  // SortOrder_ASC or SortOrder_DESC
    Cursor:   "",             // 分页 cursor
    Limit:    100,            // 默认 50
})

if err != nil {
    // handle list orders error
}

orders := results.Orders
nextCuror := results.Pagination.NextCursor
```

## 订单状态

| state | 状态 | 描述 |
|------|------|------|
| Order_Trading | 交易中 | 正在兑换 |
| Order_Filled | 已成交 | 兑换成功，将买到的币打给下单的用户 |
| Order_Cancelled | 已取消 | 取消成功，将退还币给用户 |
| Order_Rejected| 已拒绝| 下单不合法或者对应的交易所进入维护了，将退还币给用户 |
| Order_Timeout| 已超时 | 订单兑换超时，将退还币给用户；Market 单五分钟超时 |

## ATM 回款 memo

**兑换成功**
```json5
{
  "s": "ATM",
  "c": "FILLED",
  "t": "order_id"
}
```

**兑换失败**
```json5
{
  "s": "ATM",
  "c": "REFUND",
  "t": "order_id"
}
```
