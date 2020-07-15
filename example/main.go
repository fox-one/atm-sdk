package main

import (
	"net/http"

	atm "github.com/fox-one/atm-sdk/go"
	"github.com/fox-one/atm-sdk/go/proto"
)

func main() {
	markets := proto.NewMarketServiceProtobufClient(atm.Endpoint, &http.Client{})
}
