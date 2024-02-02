package blockchain

import (
	"testing"

	"github.com/blockchain-prac/internal/wallet"
	"github.com/stretchr/testify/assert"
)

func TestNewBlockchainAddBlock(t *testing.T) {
	bc := NewBlockchain()
	assert.Equal(t, 1, len(bc.Chain))
	assert.Equal(t, NewGenesisBlock(), bc.Chain[0])
}

func TestBlockchainAddBlock(t *testing.T) {
	bc := NewBlockchain()

	w := wallet.NewWallet()
	transaction, err := wallet.NewTransaction(w, "recipient", 50)
	assert.Nil(t, err)
	data := []*wallet.Transaction{transaction}

	bc.AddBlock(data)
	assert.Equal(t, data, bc.Chain[len(bc.Chain)-1].Data)
}

func TestBlockchainValidatesValidChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()

	w := wallet.NewWallet()
	transaction, err := wallet.NewTransaction(w, "recipient", 50)
	assert.Nil(t, err)
	data := []*wallet.Transaction{transaction}

	bc2.AddBlock(data)
	assert.True(t, bc.IsValidChain(bc2.Chain))
}

func TestBlockchainInvalidatesCorruptedGenesisBlock(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()

	w := wallet.NewWallet()
	transaction, err := wallet.NewTransaction(w, "recipient", 50)
	assert.Nil(t, err)

	bc2.Chain[0].Data = append(bc2.Chain[0].Data, transaction)
	assert.False(t, bc.IsValidChain(bc2.Chain))
}

func TestBlockchainInvalidatesCorruptedChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()

	w := wallet.NewWallet()
	transaction, err := wallet.NewTransaction(w, "recipient", 50)
	assert.Nil(t, err)
	data := []*wallet.Transaction{transaction}

	bc2.AddBlock(data)

	newTransaction, newErr := wallet.NewTransaction(w, "recipient", 100)
	assert.Nil(t, newErr)
	newData := []*wallet.Transaction{newTransaction}

	bc2.Chain[1].Data = newData
	assert.False(t, bc.IsValidChain(bc2.Chain))
}

func TestBlockchainReplaceValidChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()

	w := wallet.NewWallet()
	transaction, err := wallet.NewTransaction(w, "recipient", 50)
	assert.Nil(t, err)
	data := []*wallet.Transaction{transaction}

	bc2.AddBlock(data)
	bc.ReplaceChain(bc2.Chain)
	assert.Equal(t, bc2.Chain, bc.Chain)
}

func TestBlockchainCantReplaceShorterChain(t *testing.T) {
	bc := NewBlockchain()

	w := wallet.NewWallet()
	transaction, err := wallet.NewTransaction(w, "recipient", 50)
	assert.Nil(t, err)
	data := []*wallet.Transaction{transaction}

	bc.AddBlock(data)
	bc2 := NewBlockchain()
	bc.ReplaceChain(bc2.Chain)
	assert.NotEqual(t, bc2.Chain, bc.Chain)
}

func TestBlockchainCantReplaceInvalidChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()

	w := wallet.NewWallet()
	transaction, err := wallet.NewTransaction(w, "recipient", 50)
	assert.Nil(t, err)
	data := []*wallet.Transaction{transaction}

	bc2.AddBlock(data)
	bc2.Chain[1].Hash = "123"
	bc.ReplaceChain(bc2.Chain)
	assert.NotEqual(t, bc2.Chain, bc.Chain)
}
