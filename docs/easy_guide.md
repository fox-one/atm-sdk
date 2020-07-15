# Fox ATM 币币兑换简易接入指南

ATM 是 Fox.ONE 开发和维护的币币兑换引擎。引擎打通了 binance、huobi、bigone 以及 fox 自己的 f1ex 交易所；
当引擎收到用户下单之后，会去对应的交易所也下一个一样的订单；当交易所的订单成交之后，引擎会根据订单成交金额扣除手续费
之后转账给用户。如果是用 MARKET（极速兑换）策略的话，整个兑换时间一般在五秒内。ATM 是完全开放的服务，如果你的业务
有币币极速兑换的需求，可以很简单的接入使用。

完整的 api 文档可以在[这里](https://github.com/fox-one/atm-sdk/blob/master/docs/api.md)查看。如果你使用 Go 语言开发的话，
可以使用[atm-sdk](https://github.com/fox-one/atm-sdk/tree/master/go)，并且有 RPC 接口可以使用。下面我们就以 Go 语言作为
示例写一个简单的教程。

假设我们要用 USDT（ERC20) 去兑换比特币，首先调用交易对详情接口看一下 BTCUSDT 这个交易对的一些信息：

```go
package main

import (
	"context"
	"log"
	"net/http"

	atm "github.com/fox-one/atm-sdk/go"
	"github.com/fox-one/atm-sdk/go/proto"
)

func main() {
	var (
		ctx    = context.Background()
		client = &http.Client{}
		symbol = "BTCUSDT"
        userID = "your wallet id"
	)

	markets := proto.NewMarketServiceProtobufClient(atm.Endpoint, client)
	pair, err := markets.ReadPair(ctx, &proto.MarketServiceReq_ReadPair{
		Symbol: symbol,
	})
	if err != nil {
		log.Panicln(err)
	}
    
    log.Println("买单最小下单量", pair.Quote.Min)
	log.Println("买单最大下单量", pair.Quote.Max)
	log.Println("当前买一价格", pair.PriceBuy)
	log.Println("当前卖一价格", pair.PriceSell)
}
```

* 交易对详情 rest api: https://efox.fox.one/api/pairs/btcusdt
* 交易对深度 rest api: https://efox.fox.one/api/pairs/btcusdt/depth
* 交易对列表 rest api: https://efox.fox.one/api/pairs

接下来我们付 10 USDT 去下单，直接给 ATM 指定的收款钱包转账即可

```go
memo := atm.OrderAction{
    Asset:    "BTC",
    Strategy: "MARKET",
}.Encode()

putOrderTransfer := &mixin.TransferInput{
    AssetID:    pair.Quote.AssetId,
    OpponentID: atm.BrokerID,
    Amount:     decimal.NewFromInt(10),
    TraceID:    uuid.Must(uuid.NewV4()).String(),
    Memo:       memo,
}

// do putOrderTransfer
```

等待几秒钟，就会收到 ATM 转来的币了。ATM 会以一下固定的 memo 格式转账：

```json5
{
  "s": "ATM",
  "c": "REFUND or FILLED", // REFUND 退款，FILLED 兑换成功
  "t": "07975475-e713-43e8-9fe8-ea5909f73e81" // 订单 id，即下单转账的 trace id
}
``` 

这样就完成一笔兑换了。接下来更进一步，使用接口查询订单并且撤单。

首先，生成一对 rsa 密钥，将公钥放在 memo 里面转账给 ATM，私钥留在本地后续签名 token 用

```go
key, _ := rsa.GenerateKey(rand.Reader, 1024)
{
    pub := key.Public()
    data, _ := x509.MarshalPKIXPublicKey(pub)
    memo := atm.OrderAction{
        PublicKey: base64.StdEncoding.EncodeToString(data),
    }.Encode()

    registerTransfer := &mixin.TransferInput{
        AssetID:    "965e5c6e-434c-3fa9-b780-c50f43cd955c", // cnb
        OpponentID: atm.BrokerID,
        Amount:     decimal.NewFromInt(1),
        TraceID:    uuid.Must(uuid.NewV4()).String(),
        Memo:       memo,
    }

    // handle registerTransfer
}

token := atm.GenerateToken(userID, key, time.Hour)
ctx = atm.WithToken(ctx, token)

users := proto.NewUserServiceProtobufClient(atm.Endpoint, client)
if _, err := users.Me(ctx, &proto.UserServiceReq_Me{}); err == nil {
    log.Println("register ok")
}
```

**查询订单**

```go
order, err := users.ReadOrder(ctx, &proto.UserServiceReq_ReadOrder{
    TraceId: putOrderTransfer.TraceID,
})

if err != nil {
    if atm.IsErrorNotFound(err) {
        // 订单不存在。因为到账之后 ATM 那边才会创建订单，所以转账后一小段时间内查询不到是正常的
    }
}

log.Println("订单状态", order.State)
log.Println("成交金额", order.FilledAmount)
log.Println("手续费", order.FeeAmount)
```

**撤单**

```go
if _, err := users.CancelOrder(ctx, &proto.UserServiceReq_CancelOrder{
    TraceId: order.Id,
}); err == nil {
    // 撤单成功
}
```

撤单请求成功后会把订单标记成撤单状态。有时候已经标记成撤单了，但是最后订单还是成交了，是因为交易所那边的订单
在 ATM 引擎去撤单的时候已经成交了，这种情况在极速兑换策略比较常见。

**查看历史订单**

```go
orders, err := users.ListOrders(ctx, &proto.UserServiceReq_ListOrders{
	Symbol: "BTCUSDT",
})
```

以上就是 ATM 接入的简易教程，是否已经满足了你的需求呢？还不够，继续往下看。

ATM 还提供商户接入的形式：

* 商户专属的 broker 地址
* 拥有商户授权，可以查看所属的所有用户的订单和历史记录
* 可能会有的手续费分成（具体多少需要和 fox 谈）
* mixin 小群技术支持 :)

感兴趣的话，在 mixin messenger 里面联系发条 （37160854）
![发条联系方式](https://raw.githubusercontent.com/fox-one/atm-sdk/master/docs/fatiao.jpg)
