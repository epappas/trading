package binance

import "github.com/pdepip/go-binance/binance"

// DepthState Message received from depth websocket
type DepthState struct {
	EventType string          `json:"e"`
	EventTime int64           `json:"E"`
	Symbol    string          `json:"s"`
	UpdateID  int64           `json:"u"`
	BidDelta  []binance.Order `json:"b"`
	AskDelta  []binance.Order `json:"a"`
}

// KlineState Message received from kline websocket
type KlineState struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	Kline     Kline  `json:"k"`
}

// Kline Scema of KlineState
type Kline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Symbol               string `json:"s"`
	Interval             string `json:"i"`
	FirstTradeID         int32  `json:"f"`
	LastTradeID          int32  `json:"L"`
	OpenPrice            string `json:"o"`
	ClosePrice           string `json:"c"`
	HighPrice            string `json:"h"`
	LowPrice             string `json:"l"`
	Volume               string `json:"v"`
	NumberOfTrades       int32  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

// TradesState Message received from Trades websocket
type TradesState struct {
	EventType    string `json:"e"`
	EventTime    int64  `json:"E"`
	Symbol       string `json:"s"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeID int32  `json:"f"`
	LastTradeID  int32  `json:"l"`
	TradeTime    int64  `json:"T"`
	IsMaker      bool   `json:"m"`
}
