package atm

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"time"

	"github.com/fox-one/atm-sdk/go/proto"
)

const host = "https://efox.fox.one"

func ExampleAtm() {
	ctx := context.Background()

	merchantID := "merchant id"
	privateKey := "private key"

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		panic("decode pem failed: block is nil")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	token := GenerateToken(merchantID, key, time.Minute)
	ctx = WithToken(ctx, token)

	client := proto.NewMerchantServiceProtobufClient(host, &http.Client{})

	// 查询订单
	orderID := "<trace id>"
	order, err := client.ReadOrder(ctx, &proto.MerchantServiceReq_ReadOrder{
		TraceId: orderID,
	})

	if err != nil {
		if IsErrorNotFound(err) {
			// handle order not found error
		}

		// handle err
	}

	// 撤单
	if _, err := client.CancelOrder(ctx, &proto.MerchantServiceReq_CancelOrder{
		TraceId: orderID,
	}); err != nil {
		// handle cancel order failed error
	}

	// 查询历史订单
	results, err := client.ListOrders(ctx, &proto.MerchantServiceReq_ListOrders{
		Symbol:   "BTCUSDT",           // pay symbol，为空查询所有交易对
		Side:     "ASK",               // ASK or BID，为空查询所有方向
		Strategy: StrategyMarket,      // 策略，为空查询所有策略
		State:    "pending",           // pending or done，为空查询所有状态
		UserId:   order.UserId,        // 用户 id
		Order:    proto.SortOrder_ASC, // SortOrder_ASC or SortOrder_DESC
		Cursor:   "",                  // 分页 cursor
		Limit:    100,                 // 默认 50
	})

	if err != nil {
		// handle list orders error
	}

	orders := results.Orders
	nextCuror := results.Pagination.NextCursor
}
