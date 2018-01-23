package app

import (
	"sync"

	"github.com/go-kit/kit/log"

	"github.com/epappas/trading/app/backends"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// App the app main structure
type App struct {
	Version        string
	HorizonNetwork *horizon.Client
	Logger         *log.Logger
	ShapeShiftURL  string
}

type Backend interface {
	Run() error
}

var backendList []Backend

// Start the app
func (app App) Start() {
	var wg sync.WaitGroup

	backendList = append(backendList, &backends.Stellar{
		HorizonNetwork: app.HorizonNetwork,
		Logger:         app.Logger,
		Network:        build.TestNetwork,
	})

	backendList = append(backendList, &backends.ShapeShift{
		URL: app.ShapeShiftURL,
	})

	backendList = append(backendList, &backends.Binance{
		APIKEY:    "MK1gYheXyKnmIlej8U9GD5vEvED2Osb5ISipuxF2yIKR0Hq21r2w9DykUTYxgo1r",
		APISECRET: "kXgqnNCBEuxntuon9pCQLmCN1xlOL4FuEEYqRDZD3IPYIYsR268O8pdbTAy9RDYE",
		Logger:    app.Logger,
	})

	for _, backend := range backendList {
		wg.Add(1)
		go func(backend Backend) {
			defer wg.Done()
			backend.Run()
		}(backend)
	}
	wg.Wait()
}
