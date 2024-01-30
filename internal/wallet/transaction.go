package wallet

import (
	"errors"
	"fmt"

	"github.com/blockchain-prac/utils"
)

type TransactionOutput struct {
	amount  int
	address string
}
type Transaction struct {
	Id      string
	Inputs  []string
	Outputs []*TransactionOutput
}

func NewTransaction(senderWallet *Wallet, recipientAddress string, amount int) (*Transaction, error) {
	newBalance := senderWallet.Balance - amount
	if newBalance < 0 {
		return nil, errors.New(fmt.Sprintf("Amount %d exceeds wallet balance", senderWallet.Balance))
	}

	outputs := []*TransactionOutput{
		{
			amount:  newBalance,
			address: senderWallet.PublicKeyStr,
		},
		{
			amount:  amount,
			address: recipientAddress,
		},
	}

	return &Transaction{
		Id:      utils.GenerateUniqueId(),
		Inputs:  nil,
		Outputs: outputs,
	}, nil
}
