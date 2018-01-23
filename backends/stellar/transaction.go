package stellar

import (
	"container/list"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

// Transactions The managing structure
type Transactions struct {
	client  *horizon.Client
	issuer  *keypair.Full
	network build.Network
	stellar *Network
}

type OngoingTx struct {
	client  *horizon.Client
	issuer  *keypair.Full
	ops     *list.List
	stellar *Network
}

func NewTransactions(stellarNet *Network, issuer *keypair.Full) *Transactions {
	var transactions = &Transactions{
		stellarNet.client, issuer, stellarNet.network, stellarNet,
	}

	return transactions
}

func (transactions *Transactions) BeginTx() *OngoingTx {
	var ongoingTx = &OngoingTx{
		client:  transactions.client,
		issuer:  transactions.issuer,
		ops:     list.New(),
		stellar: transactions.stellar,
	}

	ongoingTx.ops.PushFront(build.SourceAccount{AddressOrSeed: transactions.issuer.Address()})
	ongoingTx.ops.PushFront(transactions.network)
	ongoingTx.ops.PushFront(build.AutoSequence{SequenceProvider: transactions.client})

	return ongoingTx
}

func (transactions *Transactions) Asset(code string) build.Asset {
	if len(code) > 0 {
		return build.CreditAsset(code, transactions.issuer.Address())
	}
	return build.NativeAsset()
}

func (transactions *Transactions) OtherAsset(code, address string) build.Asset {
	if len(code) > 0 {
		return build.CreditAsset(code, address)
	}
	return build.NativeAsset()
}

func (transactions *Transactions) BeginPathPayment(amount string, code string) build.PayWithPath {
	return build.PayWith(transactions.Asset(code), amount)
}

func (ongoingTx *OngoingTx) AddPayment(receiver keypair.KP, amount string) *OngoingTx {
	ongoingTx.ops.PushFront(build.Payment(
		build.Destination{AddressOrSeed: receiver.Address()},
		build.NativeAmount{Amount: amount},
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) AddAssetPayment(receiver keypair.KP, code string, amount string) *OngoingTx {
	ongoingTx.ops.PushFront(build.Payment(
		build.Destination{AddressOrSeed: receiver.Address()},
		build.CreditAmount{
			Code:   code,
			Issuer: ongoingTx.issuer.Address(),
			Amount: amount,
		},
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) AddPathPayment(receiver keypair.KP, code string, amount string, payPath *build.PayWithPath) *OngoingTx {
	ongoingTx.ops.PushFront(build.Payment(
		build.Destination{AddressOrSeed: receiver.Address()},
		build.CreditAmount{
			Code:   code,
			Issuer: ongoingTx.issuer.Address(),
			Amount: amount,
		},
		payPath,
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) CreateOffer(rate build.Rate, amount string) *OngoingTx {
	ongoingTx.ops.PushFront(build.CreateOffer(rate, build.Amount(amount)))

	return ongoingTx
}

func (ongoingTx *OngoingTx) CreatePassiveOffer(rate build.Rate, amount string) *OngoingTx {
	ongoingTx.ops.PushFront(build.CreatePassiveOffer(rate, build.Amount(amount)))

	return ongoingTx
}

func (ongoingTx *OngoingTx) DeleteOffer(rate build.Rate, offerID uint64) *OngoingTx {
	ongoingTx.ops.PushFront(build.DeleteOffer(rate, build.OfferID(offerID)))

	return ongoingTx
}

func (ongoingTx *OngoingTx) UpdateOffer(rate build.Rate, amount build.Amount, offerID uint64) *OngoingTx {
	ongoingTx.ops.PushFront(build.UpdateOffer(rate, amount, build.OfferID(offerID)))

	return ongoingTx
}

func (ongoingTx *OngoingTx) ChangeTrust(receiver *keypair.Full, code string, limit string) *OngoingTx {
	ongoingTx.ops.PushFront(build.ChangeTrust(
		build.CreditAsset(code, ongoingTx.issuer.Address()),
		build.Limit(limit),
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) RemoveTrust(receiver *keypair.Full, code string, limit string) *OngoingTx {
	ongoingTx.ops.PushFront(build.RemoveTrust(
		code,
		ongoingTx.issuer.Address(),
		build.SourceAccount{AddressOrSeed: receiver.Address()},
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) AllowTrust(receiver *keypair.Full, code string, allow bool) *OngoingTx {
	ongoingTx.ops.PushFront(build.AllowTrust(
		build.Trustor{Address: receiver.Address()},
		build.AllowTrustAsset{Code: code},
		build.Authorize{Value: allow},
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) CreateAccount(account keypair.KP, amount string) *OngoingTx {
	ongoingTx.ops.PushFront(build.CreateAccount(
		build.Destination{AddressOrSeed: account.Address()},
		build.NativeAmount{Amount: amount},
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) AccountMerge(account keypair.KP) *OngoingTx {
	ongoingTx.ops.PushFront(build.AccountMerge(
		build.Destination{AddressOrSeed: account.Address()},
	))

	return ongoingTx
}

func (ongoingTx *OngoingTx) SetOptions(muts ...interface{}) *OngoingTx {
	ongoingTx.ops.PushFront(build.SetOptions(muts...))

	return ongoingTx
}

func (ongoingTx *OngoingTx) Inflation() *OngoingTx {
	ongoingTx.ops.PushFront(build.Inflation())

	return ongoingTx
}

func (ongoingTx *OngoingTx) SubmitTx(seeds ...string) (*horizon.TransactionSuccess, error) {
	count := ongoingTx.ops.Len()
	ops := make([]build.TransactionMutator, 0, count+3)

	seeds = append(seeds, ongoingTx.issuer.Seed())

	for tx := ongoingTx.ops.Back(); tx != nil; tx = tx.Prev() {
		ops = append(ops, tx.Value.(build.TransactionMutator))
	}

	txBuilder := build.Transaction(ops...)
	success, err := ongoingTx.stellar.SubmitTx(txBuilder, seeds...)
	if err != nil {
		return success, err
	}

	return success, nil
}
