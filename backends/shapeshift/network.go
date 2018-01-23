package shapeshift

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiURL = "https://shapeshift.io"

type Network struct {
	URL string
}

func NewNetwork(url string) *Network {
	if len(url) == 0 {
		return &Network{apiURL}
	}

	return &Network{url}
}

type HTTP_METHOD interface {
	NewRequest(data interface{}) (*http.Request, error)
}

type POST struct {
	endpoint string
	network  *Network
}

type GET struct {
	endpoint string
	network  *Network
}

func (post POST) NewRequest(data interface{}) (*http.Request, error) {
	var url = fmt.Sprintf("%s/%s", post.network.URL, post.endpoint)
	new, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(new))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func (get GET) NewRequest(query interface{}) (*http.Request, error) {
	var url = fmt.Sprintf("%s/%s/%s", get.network.URL, get.endpoint, query)
	return http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
}

func (network *Network) DoHttp(method HTTP_METHOD, data interface{}) ([]byte, error) {
	req, err := method.NewRequest(data)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}

func (network *Network) HTTPCall(method HTTP_METHOD, data interface{}, schema interface{}) error {
	r, err := network.DoHttp(method, data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(r, schema)
	return err
}

func (network *Network) RecentTransactions(count string) (*RecentTranxResponse, error) {
	var resp RecentTranxResponse
	var method = &GET{"recenttx", network}
	err := network.HTTPCall(method, count, &resp)

	return &resp, err
}

func (network *Network) DepositStatus(addr string) (*DepositStatusResponse, error) {
	var resp DepositStatusResponse
	var method = &GET{"txStat", network}
	err := network.HTTPCall(method, addr, &resp)

	return &resp, err
}

func (network *Network) TimeRemaining(addr string) (*TimeRemainingResponse, error) {
	var resp TimeRemainingResponse
	var method = &GET{"timeremaining", network}
	err := network.HTTPCall(method, addr, &resp)

	return &resp, err
}

func (network *Network) Coins(addr string) (*CoinsResponse, error) {
	var resp CoinsResponse
	var method = &GET{"getcoins", network}
	err := network.HTTPCall(method, addr, &resp)

	return &resp, err
}

func (network *Network) Validate(addr string, coin string) (*ValidateResponse, error) {
	var resp ValidateResponse
	var method = &GET{"validateAddress/" + addr, network}
	err := network.HTTPCall(method, coin, &resp)

	return &resp, err
}

func (network *Network) Cancel(addr Address) (*CancelResponse, error) {
	var resp CancelResponse
	var method = &POST{"cancelpending", network}
	err := network.HTTPCall(method, addr, &resp)

	return &resp, err
}

func (network *Network) SendReceipt(receipt Receipt) (*ReceiptResponse, error) {
	var resp ReceiptResponse
	var method = &POST{"mail", network}
	err := network.HTTPCall(method, receipt, &resp)

	return &resp, err
}

func (network *Network) Shift(newPair NewPair) (*NewTransactionResponse, error) {
	var resp NewTransactionResponse
	var method = &POST{"shift", network}
	err := network.HTTPCall(method, newPair, &resp)

	return &resp, err
}

func (network *Network) SendAmount(newPair NewPair) (*NewFixedTransactionResponse, error) {
	var resp NewFixedTransactionResponse
	var method = &POST{"sendamount", network}
	err := network.HTTPCall(method, newPair, &resp)

	return &resp, err
}

func (network *Network) ListTxByAddress(addr string, key string) (*[]Transaction, error) {
	var resp []Transaction
	var method = &GET{"txbyaddress" + addr, network}
	err := network.HTTPCall(method, key, &resp)

	return &resp, err
}

func (network *Network) ListTxByKey(key string) (*[]Transaction, error) {
	var resp []Transaction
	var method = &GET{"txbyapikey", network}
	err := network.HTTPCall(method, key, &resp)

	return &resp, err
}

func (network *Network) GetRates(p Pair) (float64, error) {
	var resp RateResponse
	var method = &GET{"txbyapikey", network}
	err := network.HTTPCall(method, p.Name, &resp)

	return ToFloat(resp.Rate), err
}

func (network *Network) GetLimit(p Pair) (float64, error) {
	var resp LimitResponse
	var method = &GET{"limit", network}
	err := network.HTTPCall(method, p.Name, &resp)

	return ToFloat(resp.Limit), err
}

func (network *Network) GetInfo(p Pair) (*MarketInfoResponse, error) {
	var resp MarketInfoResponse
	var method = &GET{"marketinfo", network}
	err := network.HTTPCall(method, p.Name, &resp)

	return &resp, err
}
