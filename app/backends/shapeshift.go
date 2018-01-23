package backends

import (
	"flag"
	"fmt"
	"time"

	backend "github.com/epappas/trading/backends/shapeshift"
)

type ShapeShift struct {
	URL string
}

func (shapeShift *ShapeShift) Run() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[shapeShift] Recovered", r)
		}
	}()

	flag.Parse()

	fmt.Println("Inside ShapeShift Here")

	var network = backend.NewNetwork("")

	var resp, err = network.SendAmount(backend.NewPair{
		Pair:        "btc_eth",
		Amount:      float64(0.5),
		ToAddress:   "0xCD0728052b45bE2a66695F5352F39a0184066461",
		FromAddress: "1M8yWFAQskbdATLG3kPwMgcr1Y54zE8ijM",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("ShapeShift Shift response", resp)

	time.Sleep(time.Second * 5)

	var time, err2 = network.TimeRemaining(resp.Withdrawal)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("ShapeShift TimeRemaining response", time)

	return nil
}
