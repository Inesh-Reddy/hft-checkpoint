package main

import (
	"context"
	"fmt"

	"github.com/Inesh-Reddy/hft-checkpoint/packages/golib/redis"
)



func main(){
	fmt.Println("Ticker Service running .....");
	r:=redis.NewClient()
	r.Set(context.Background(), "binance.btcusdt.ticker", "Inesh")
	
}