# Fox ATM 币币兑换简易接入指南

ATM 是 Fox.ONE 开发和维护的币币兑换引擎。引擎打通了 binance、huobi、bigone 以及 fox 自己的 f1ex 交易所；
当引擎收到用户下单之后，会去对应的交易所也下一个一样的订单；当交易所的订单成交之后，引擎会根据订单成交金额扣除手续费
之后转账给用户。如果是用 MARKET（极速兑换）策略的话，整个兑换时间一般在五秒内。ATM 是完全开放的服务，如果你的业务
有币币极速兑换的需求，可以很简单的接入使用。

完整的 api 文档可以在[这里](https://github.com/fox-one/atm-sdk/docs/api.md)查看。如果你使用 Go 语言开发的话，
可以使用[atm-sdk](https://github.com/fox-one/atm-sdk/go)，并且有 RPC 接口可以使用。下面我们就以 Go 语言作为
示例写一个简单的教程。

假设我们要用 USDT（ERC20) 去兑换比特币，首先调交易对详情接口看一下 BTCUSDT 这个交易对的一些信息：

```go
package main
```





