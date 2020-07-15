package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	atm "github.com/fox-one/atm-sdk/go"
	"github.com/fox-one/atm-sdk/go/proto"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
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

	// 下单
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

	// handle putOrderTransfer

	// register public key
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

	// 查询订单
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

	// 撤单
	if _, err := users.CancelOrder(ctx, &proto.UserServiceReq_CancelOrder{
		TraceId: order.Id,
	}); err == nil {
		// 撤单成功
	}

	// 历史订单
	orders, err := users.ListOrders(ctx, &proto.UserServiceReq_ListOrders{
		Symbol: "BTCUSDT",
	})

}
