package binance

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/pdepip/go-binance/binance"
)

type Network struct {
	apiURL    string
	wsURL     string
	apiKEY    string
	apiSECRET string
	Client    *binance.Binance
}

func NewNetwork(apiKey, apiSecret string) *Network {
	return &Network{
		"https://api.binance.com",
		"wss://stream.binance.com:9443",
		apiKey, apiSecret,
		binance.New(apiKey, apiSecret),
	}
}

func (network *Network) WSStream(address string, handler func([]byte)) {
	var wsDialer websocket.Dialer
	wsConn, _, err := wsDialer.Dial(address, nil)
	if err != nil {
		panic(err)
	}
	defer wsConn.Close()

	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("[ERROR] WSStream ReadMessage:", address, err)
			continue
		}

		handler(message)
	}
}

func (network *Network) StreamDepth(symbol string, stateChan chan DepthState) {
	var address = fmt.Sprintf("%s/ws/%s@depth", network.wsURL, symbol)

	network.WSStream(address, func(message []byte) {
		msg := DepthState{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("[ERROR] StreamDepth Parsing:", err)
		} else {
			stateChan <- msg
		}
	})
}

func (network *Network) StreamKline(symbol, interval string, stateChan chan KlineState) {
	var address = fmt.Sprintf("%s/ws/%s@kline_%s", network.wsURL, symbol, interval)

	network.WSStream(address, func(message []byte) {
		msg := KlineState{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("[ERROR] StreamKline Parsing:", err)
		} else {
			stateChan <- msg
		}
	})
}

func (network *Network) StreamTrades(symbol string, stateChan chan TradesState) {
	var address = fmt.Sprintf("%s/ws/%s@aggTrade", network.wsURL, symbol)

	network.WSStream(address, func(message []byte) {
		msg := TradesState{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("[ERROR] StreamTrades Parsing:", err)
		} else {
			stateChan <- msg
		}
	})
}
