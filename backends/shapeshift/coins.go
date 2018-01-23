package shapeshift

type Coin struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	Image           string `json:"image,omitempty"`
	SpecialReturn   bool   `json:"specialReturn,omitempty"`
	SpecialOutgoing bool   `json:"specialOutgoing,omitempty"`
	SpecialIncoming bool   `json:"specialIncoming,omitempty"`
	FieldName       string `json:"fieldName,omitempty"`
	FieldKey        string `json:"fieldKey,omitempty"`
	QrName          string `json:"qrName,omitempty"`
}

type CryptoCoin struct {
	Coin
}

type CoinsResponse struct {
	BTC   CryptoCoin `json:"BTC"`
	BCY   CryptoCoin `json:"BCY"`
	BLK   CryptoCoin `json:"BLK"`
	BTCD  CryptoCoin `json:"BTCD"`
	BTS   CryptoCoin `json:"BTS"`
	CLAM  CryptoCoin `json:"CLAM"`
	DASH  CryptoCoin `json:"DASH"`
	DGB   CryptoCoin `json:"DGB"`
	DGD   CryptoCoin `json:"DGD"`
	DOGE  CryptoCoin `json:"DOGE"`
	EMC   CryptoCoin `json:"EMC"`
	ETH   CryptoCoin `json:"ETH"`
	ETC   CryptoCoin `json:"ETC"`
	FCT   CryptoCoin `json:"FCT"`
	GNT   CryptoCoin `json:"GNT"`
	LBC   CryptoCoin `json:"LBC"`
	LSK   CryptoCoin `json:"LSK"`
	LTC   CryptoCoin `json:"LTC"`
	MAID  CryptoCoin `json:"MAID"`
	MONA  CryptoCoin `json:"MONA"`
	MSC   CryptoCoin `json:"MSC"`
	NBT   CryptoCoin `json:"NBT"`
	NMC   CryptoCoin `json:"NMC"`
	NVC   CryptoCoin `json:"NVC"`
	NXT   CryptoCoin `json:"NXT"`
	POT   CryptoCoin `json:"POT"`
	PPC   CryptoCoin `json:"PPC"`
	REP   CryptoCoin `json:"REP"`
	RDD   CryptoCoin `json:"RDD"`
	SDC   CryptoCoin `json:"SDC"`
	SC    CryptoCoin `json:"SC"`
	SJCX  CryptoCoin `json:"SJCX"`
	START CryptoCoin `json:"START"`
	STEEM CryptoCoin `json:"STEEM"`
	SNGLS CryptoCoin `json:"SNGLS"`
	USDT  CryptoCoin `json:"USDT"`
	VOX   CryptoCoin `json:"VOX"`
	VRC   CryptoCoin `json:"VRC"`
	VTC   CryptoCoin `json:"VTC"`
	XCP   CryptoCoin `json:"XCP"`
	XMR   CryptoCoin `json:"XMR"`
	XRP   CryptoCoin `json:"XRP"`
	ZEC   CryptoCoin `json:"ZEC"`
}
