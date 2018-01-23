package backends

import (
	"flag"
	"fmt"

	"github.com/go-kit/kit/log"

	backend "github.com/epappas/trading/backends/binance"
)

type Binance struct {
	APIKEY    string
	APISECRET string
	Logger    *log.Logger
	Network   *backend.Network
}

func (binance *Binance) Run() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[Binance] Recovered", r)
		}
	}()

	flag.Parse()

	fmt.Println("Inside Binance Here")

	binance.Network = backend.NewNetwork(binance.APIKEY, binance.APISECRET)
	var aggregateChan = make(chan string, 100)

	go (func() {
		var streamChan = make(chan backend.DepthState, 10)

		go binance.Network.StreamDepth("ethbtc", streamChan)
		for state := range streamChan {
			aggregateChan <- fmt.Sprintf("Binance DepthState %s", state)
		}
	})()

	go (func() {
		var streamChan = make(chan backend.KlineState, 10)

		go binance.Network.StreamKline("ethbtc", "1m", streamChan)
		for state := range streamChan {
			aggregateChan <- fmt.Sprintf("Binance KlineState %s", state)
		}
	})()

	go (func() {
		var streamChan = make(chan backend.TradesState, 10)

		go binance.Network.StreamTrades("ethbtc", streamChan)
		for state := range streamChan {
			aggregateChan <- fmt.Sprintf("Binance TradesState %s", state)
		}
	})()

	for msg := range aggregateChan {
		fmt.Println(msg)
	}

	return nil
}
