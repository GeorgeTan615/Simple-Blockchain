package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountDeductedFromWalletBalance(t *testing.T) {
	wallet := NewWallet()
	amount := 50
	recipientAddress := "recipient"
	transaction, err := NewTransaction(wallet, recipientAddress, amount)

	assert.Nil(t, err)

	for _, output := range transaction.Outputs {
		if output.address == wallet.PublicKeyStr {
			assert.Equal(t, wallet.Balance-amount, output.amount)
		}

		if output.address == recipientAddress {
			assert.Equal(t, amount, output.amount)
		}
	}
}

func TestAmountExceedsWalletBalance(t *testing.T) {
	wallet := NewWallet()
	amount := 500000
	recipientAddress := "recipient"
	transaction, err := NewTransaction(wallet, recipientAddress, amount)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
}
