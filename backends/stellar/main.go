package stellar

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

// Genesis Generate a Key pair
func Genesis() (kp *keypair.Full) {
	pair, err := keypair.Random()
	if err != nil {
		panic(err)
	}

	return pair
}

// CreateAccount a single account and fund it some money
func CreateAccount(txBuilder *OngoingTx, amount string) keypair.KP {
	kp := Genesis()

	txBuilder.CreateAccount(kp, amount)

	return kp
}

func LogBalance(account *horizon.Account, logger log.Logger) {
	for _, balance := range account.Balances {
		level.Info(logger).Log("balance", balance.Balance, "asset_type", balance.Asset.Type)
	}
}

func LogBalances(network *Network, keypairs []keypair.KP, logger log.Logger) {
	for i, kp := range keypairs {
		l := log.With(logger, "account_index", i)
		if kp != nil {
			acc, err := network.LoadAccount(kp.Address())
			if err != nil {
				panic(err)
			}
			LogBalance(acc, l)
		}
	}
}
