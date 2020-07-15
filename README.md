# atm-sdk
fox atm sdk

* api host `https://efox.fox.one`
* atm 默认收款地址 fb151d94-cc0e-358b-9cdb-57215a557a31 
* atm 注册商户有自己单独的收款地址

## 注册

* 生成 rsa 公私钥对，将公钥放在 memo 转账给 atm
* 后续用这个私钥作为 token 的签名私钥

**Memo**

```go
key, _ := rsa.GenerateKey(rand.Reader, 1024)
pub := key.Public()
data,_ := x509.MarshalPKIXPublicKey(pub)
action := atm.OrderAction{
    PublicKey: base64.StdEncoding.EncodeToString(data),
}
memo := action.Encode()

// 用上面生成的 memo 随便转一笔 cnb 之类的不值钱的币给 atm 即可完成注册
```

## 下单

* 构造下单 memo，带上想买的币的 symbol/asset_id 和所使用的策略
* 转账付款的币给指定下单地址
* 转账的 `TraceID` 即这笔订单的订单 id，转账人的钱包 id 为这笔订单的用户 id

**Memo**

```go
action := atm.OrderAction{
    Asset: "BTC", // symbol or asset id
    Strategy: atm.StrategyMarket, // 默认是 StrategyMarket
}
memo := action.Encode()

// 用上面生成的 memo 转账你想花费的币给 atm 即可完成下单
```

## 查询订单

下单之后可以开始用下单转账的 `TraceID` 作为订单 id 查询订单。
因为异步到账的原因，可能会在一小段时间内查询不到订单

```go
client := proto.NewMerchantServiceProtobufClient(atm.Endpoint, &http.Client{})

token := atm.GenerateToken(merchantID, key, time.Minute)
ctx := atm.WithToken(context.Backgroud(), token)

orderID := "<trace id>"
order, err := client.ReadOrder(ctx, &proto.MerchantServiceReq_ReadOrder{
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
token := atm.GenerateToken(merchantID, key, time.Minute)
ctx := atm.WithToken(context.Backgroud(), token)

if _, err := client.CancelOrder(ctx, &proto.MerchantServiceReq_CancelOrder{
    TraceId: orderID,
}); err != nil {
    // handle cancel order failed error
}
```

## 查询用户历史订单

```go
token := atm.GenerateToken(merchantID, key, time.Minute)
ctx := atm.WithToken(context.Backgroud(), token)

results, err := client.ListOrders(ctx, &proto.MerchantServiceReq_ListOrders{
    Symbol:   "BTCUSDT",      // pay symbol，为空查询所有交易对
    Side:     "ASK",          // ASK or BID，为空查询所有方向
    Strategy: StrategyMarket, // 策略，为空查询所有策略
    State:    "pending",      // pending or done，为空则查询所有状态
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
