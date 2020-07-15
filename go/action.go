package atm

import (
	"encoding/json"
)

const (
	// 市价极速兑换
	StrategyMarket = "MARKET"
	// 跟随买一卖一下单
	StrategyFollow = "FOLLOW"
)

type OrderAction struct {
	// Asset 为要换取的币的 symbol 或者 mixin asset id
	Asset string `json:"a,omitempty"`
	// MARKET or FOLLOW
	Strategy string `json:"s,omitempty"`
	// 注册或者更新 public key, rsa public key base64 格式
	PublicKey string `json:"u,omitempty"`
}

func (action OrderAction) Encode() string {
	b, _ := json.Marshal(action)
	return string(b)
}

const (
	// 兑换失败，退款
	TransferActionSourceRefund = "REFUND"
	// 兑换成功
	TransferActionSourceFilled = "FILLED"
)

// TransferAction 是 atm 兑换完成或者失败后回款的 memo 格式
type TransferAction struct {
	// Service = ATM
	Service string `json:"s,omitempty"`
	// REFUND or FILLED
	Source string `json:"c,omitempty"`
	// 订单 id
	OrderID string `json:"t,omitempty"`
}
