package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	app "github.com/epappas/trading/app"
	"github.com/stellar/go/clients/horizon"
)

var (
	horizon_addr = flag.String("horizon", "http://horizon-testnet.stellar.org", "horizon address")
)

func init() {
	flag.Parse()
}

func main() {
	var logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = level.NewFilter(logger, level.AllowDebug())
	logger = log.With(logger, "time", log.DefaultTimestampUTC())
	logger = log.With(logger, "caller", log.Caller(3))

	var horizonNetwork = &horizon.Client{
		URL:  *horizon_addr,
		HTTP: http.DefaultClient,
	}

	var thisApp = app.App{
		HorizonNetwork: horizonNetwork,
		Logger:         &logger,
	}

	thisApp.Start()
	// routes.Start()
}
