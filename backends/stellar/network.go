package stellar

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

// Network The bridge for the Stellar Network Operations
type Network struct {
	client  *horizon.Client
	network build.Network
}

type Stream struct {
	network *Network
	ctx     *context.Context
	cancel  *context.CancelFunc
	cursor  *horizon.Cursor
}

// NewNetwork Create a bridge network for the given Stellar endpoint
func NewNetwork(client *horizon.Client, network build.Network) *Network {
	return &Network{client, network}
}

// Transactions generate Transactions instance out of Accounts' config
func (network *Network) Transactions(issuer *keypair.Full) *Transactions {
	return NewTransactions(network, issuer)
}

// Stream initilise the Streams API
func (network *Network) Stream(cursor string) *Stream {
	var ctx, cancel = context.WithCancel(context.Background())
	var c = horizon.Cursor(cursor)

	return &Stream{network, &ctx, &cancel, &c}
}

// HomeDomainForAccount load the home domain of an account
func (network *Network) HomeDomainForAccount(address string) (string, error) {
	domain, err := network.client.HomeDomainForAccount(address)
	if err != nil {
		return "", err
	}

	return domain, nil
}

// LoadAccount an account from the Net
func (network *Network) LoadAccount(address string) (*horizon.Account, error) {
	account, err := network.client.LoadAccount(address)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// LoadAccountOffers an account from the Net
func (network *Network) LoadAccountOffers(address string, params ...interface{}) (*horizon.OffersPage, error) {
	offers, err := network.client.LoadAccountOffers(address, params...)
	if err != nil {
		return nil, err
	}

	return &offers, nil
}

// LoadMemo a payment's memo
func (network *Network) LoadMemo(payment *horizon.Payment) (*horizon.Payment, error) {
	err := network.client.LoadMemo(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// LoadOrderBook loads order book for given selling and buying assets.
func (network *Network) LoadOrderBook(selling horizon.Asset, buying horizon.Asset, params ...interface{}) (orderBook horizon.OrderBookSummary, err error) {
	return network.client.LoadOrderBook(selling, buying, params...)
}

// SubmitTx Transmit the given Tx to the stellar Network
func (network *Network) SubmitTx(txBuilder *build.TransactionBuilder, seeds ...string) (*horizon.TransactionSuccess, error) {
	var success horizon.TransactionSuccess
	var err error

	txEnv := txBuilder.Sign(seeds...)
	txEnvB64, err := txEnv.Base64()
	if err != nil {
		return nil, err
	}

	success, err = network.client.SubmitTransaction(txEnvB64)
	if err != nil {
		return &success, err
	}
	return &success, nil
}

// FriendBot Use the testNet service to fund an account
func (network *Network) FriendBot(address string) ([]byte, error) {
	res, err := network.client.HTTP.Get(fmt.Sprintf("%s/friendbot?addr=%s", network.client.URL, address))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusBadRequest {
		return nil, errors.New("funding failure")
	}

	return data, nil
}

func (stream *Stream) Payments(accountID string, handler horizon.PaymentHandler) error {
	return stream.network.client.StreamPayments(*stream.ctx, accountID, stream.cursor, handler)
}

func (stream *Stream) AccountTx(accountID string, handler horizon.TransactionHandler) error {
	return stream.network.client.StreamTransactions(*stream.ctx, accountID, stream.cursor, handler)
}

func (stream *Stream) Ledgers(handler horizon.LedgerHandler) error {
	return stream.network.client.StreamLedgers(*stream.ctx, stream.cursor, handler)
}

func (stream *Stream) Transactions(handler horizon.TransactionHandler) error {
	return stream.network.client.StreamAllTransactions(*stream.ctx, stream.cursor, handler)
}
