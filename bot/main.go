package main

import (
	"fmt"
)

func testExchange(e Exchange) {
	res := e.Status()
	if res == Running {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}

func main() {
	exchange := BinanceExchange{Endpoint: "https://testnet.binance.vision"}
	testExchange(exchange)

	fmt.Println(exchange.Assets())
}
