package shapeshift

type Pair struct {
	Name string `json:"pair,omitempty"`
}

type RateResponse struct {
	Pair string `json:"pair,omitempty"`
	Rate string `json:"rate"`
	Error
}

type LimitResponse struct {
	Pair  string `json:"pair,omitempty"`
	Limit string `json:"limit"`
	Error
}

type MarketInfoResponse struct {
	Pair     string  `json:"pair,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
	Limit    float64 `json:"limit,omitempty"`
	Min      float64 `json:"min,omitempty"`
	MinerFee float64 `json:"minerFee,omitempty"`
	Error
}

type RecentTranxResponse []struct {
	CurIn     string  `json:"curIn"`
	CurOut    string  `json:"curOut"`
	Timestamp float64 `json:"timestamp"`
	Amount    float64 `json:"amount"`
	Error
}

type DepositStatusResponse struct {
	Status       string  `json:"status"`
	Address      string  `json:"address"`
	Withdraw     string  `json:"withdraw,omitempty"`
	IncomingCoin float64 `json:"incomingCoin,omitempty"`
	IncomingType string  `json:"incomingType,omitempty"`
	OutgoingCoin string  `json:"outgoingCoin,omitempty"`
	OutgoingType string  `json:"outgoingType,omitempty"`
	Transaction  string  `json:"transaction,omitempty"`
	Error
}

type Receipt struct {
	Email         string `json:"email"`
	TransactionID string `json:"txid"`
}

type ValidateResponse struct {
	Valid bool `json:"isValid"`
	Error
}

type CancelResponse struct {
	Success string `json:"success,omitempty"`
	Error
}

type Address struct {
	ID string `json:"address"`
}

type NewPair struct {
	Pair        string  `json:"pair,omitempty"`
	ToAddress   string  `json:"withdrawal"`
	FromAddress string  `json:"returnAddress,omitempty"`
	DestTag     string  `json:"destTag,omitempty"`
	RsAddress   string  `json:"rsAddress,omitempty"`
	APIKey      string  `json:"apiKey,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}

type NewTransactionResponse struct {
	SendTo     string `json:"deposit"`
	SendType   string `json:"depositType"`
	ReturnTo   string `json:"withdrawal"`
	ReturnType string `json:"withdrawalType"`
	Public     string `json:"public"`
	XrpDestTag string `json:"xrpDestTag"`
	APIKey     string `json:"apiPubKey"`
	Error
}

type NewFixedTransactionResponse struct {
	OrderID          string  `json:"orderId"`
	Pair             string  `json:"pair,omitempty"`
	Withdrawal       string  `json:"withdrawal"`
	WithdrawalAmount string  `json:"withdrawalAmount"`
	Deposit          string  `json:"deposit"`
	DepositAmount    string  `json:"depositAmount"`
	Expiration       int64   `json:"expiration"`
	QuotedRate       string  `json:"quotedRate"`
	MaxLimit         float64 `json:"maxLimit"`
	ReturnAddress    string  `json:"returnAddress"`
	APIPubKey        string  `json:"apiPubKey"`
	MinerFee         string  `json:"minerFee"`
	Error
}

type Transaction struct {
	InputTXID      string  `json:"inputTXID"`
	InputAddress   string  `json:"inputAddress"`
	InputCurrency  string  `json:"inputCurrency,omitempty"`
	InputAmount    float64 `json:"inputAmount,omitempty"`
	OutputTXID     string  `json:"outputTXID,omitempty"`
	OutputAddress  string  `json:"outputAddress,omitempty"`
	OutputCurrency string  `json:"outputCurrency,omitempty"`
	OutputAmount   string  `json:"outputAmount,omitempty"`
	ShiftRate      string  `json:"shiftRate,omitempty"`
	Status         string  `json:"status,omitempty"`
}

type ReceiptResponse struct {
	Email struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"email"`
	Error
}

type TimeRemainingResponse struct {
	Status  string `json:"status"`
	Seconds string `json:"seconds_remaining"`
	Error
}
