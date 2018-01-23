package binance

import (
	"fmt"
	"sync"

	"github.com/pdepip/go-binance/binance"
)

// OrderBook OrderBook structure
type OrderBook struct {
	Symbol string

	Bids     map[float64]float64
	BidMutex sync.Mutex

	Asks     map[float64]float64
	AskMutex sync.Mutex

	DepthUpdates chan DepthState
}

func NewOrderBook(symbol string, maxDepth int8) *OrderBook {
	orderBook := OrderBook{Symbol: symbol}
	orderBook.Bids = make(map[float64]float64, maxDepth)
	orderBook.Asks = make(map[float64]float64, maxDepth)
	orderBook.DepthUpdates = make(chan DepthState, 500)

	return &orderBook
}

// FetchOrders incoming orders from a network
func (ob *OrderBook) FetchOrders(network *Network) error {
	query := binance.OrderBookQuery{
		Symbol: ob.Symbol,
	}

	orderBook, err := network.Client.GetOrderBook(query)
	if err != nil {
		return err
	}

	ob.ProcessBids(orderBook.Bids)
	ob.ProcessAsks(orderBook.Asks)

	return nil
}

// ProcessBids Process all incoming bids
func (ob *OrderBook) ProcessBids(bids []binance.Order) {
	for _, bid := range bids {
		ob.BidMutex.Lock()
		if bid.Quantity == 0 {
			delete(ob.Bids, bid.Price)
		} else {
			ob.Bids[bid.Price] = bid.Quantity
		}
		fmt.Println("Bids", ob.Bids)
		ob.BidMutex.Unlock()
	}
}

// ProcessAsks Process all incoming asks
func (ob *OrderBook) ProcessAsks(asks []binance.Order) {
	for _, ask := range asks {
		ob.AskMutex.Lock()
		if ask.Quantity == 0 {
			delete(ob.Asks, ask.Price)
		} else {
			ob.Asks[ask.Price] = ask.Quantity
		}
		fmt.Println("Asks", ob.Asks)
		ob.AskMutex.Unlock()
	}
}

// Start Hands off incoming messages to processing functions
func (ob *OrderBook) Start() {
	for {
		select {
		case job := <-ob.DepthUpdates:
			if len(job.BidDelta) > 0 {
				go ob.ProcessBids(job.BidDelta)
			}

			if len(job.AskDelta) > 0 {
				go ob.ProcessAsks(job.AskDelta)
			}
		}
	}
}
