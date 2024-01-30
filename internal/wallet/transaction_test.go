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
		if output.Address == wallet.PublicKeyStr {
			assert.Equal(t, wallet.Balance-amount, output.Amount)
		}

		if output.Address == recipientAddress {
			assert.Equal(t, amount, output.Amount)
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

func TestInputHasBeenCreatedInTransaction(t *testing.T) {
	wallet := NewWallet()
	recipientAddress := "recipient"
	transaction, err := NewTransaction(wallet, recipientAddress, 0)
	assert.Nil(t, err)
	assert.Equal(t, wallet.Balance, transaction.Input.Amount)
}

func TestVerifyTransaction(t *testing.T) {
	// Happy Path
	wallet := NewWallet()
	recipientAddress := "recipient"
	transaction, err := NewTransaction(wallet, recipientAddress, 0)
	assert.Nil(t, err)
	assert.True(t, VerifyTransaction(transaction))

	// Transaction sender public key gets tampered
	transaction.Input.Address = "random"
	assert.False(t, VerifyTransaction(transaction))

	// Someone tampers with data
	transaction.Outputs[0].Amount = 100000
	assert.False(t, VerifyTransaction(transaction))
}

func TestUpdateTransaction(t *testing.T) {
	// Happy Path
	wallet := NewWallet()
	recipientAddress := "recipient"
	firstTransactionAmount := 10
	transaction, err := NewTransaction(wallet, recipientAddress, firstTransactionAmount)
	assert.Nil(t, err)

	secondTransactionAmount := 20
	err = transaction.Update(wallet, "newRecipient", secondTransactionAmount)

	senderOutput := transaction.Outputs[0]
	assert.Nil(t, err)
	assert.Equal(t, 3, len(transaction.Outputs))
	assert.Equal(
		t,
		wallet.Balance-firstTransactionAmount-secondTransactionAmount,
		senderOutput.Amount)
	assert.True(t, VerifyTransaction(transaction))

}
