package backends

import (
	"flag"
	"fmt"

	"github.com/go-kit/kit/log"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"

	backend "github.com/epappas/trading/backends/stellar"
)

type Stellar struct {
	HorizonNetwork *horizon.Client
	Logger         *log.Logger
	Network        build.Network
}

func (stellar *Stellar) Run() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[Stellar] Recovered", r)
			backend.GetTxErrorResultCodes(r.(error), *stellar.Logger)
		}
	}()

	flag.Parse()

	fmt.Println("Inside Stellar Here")

	return nil
}
