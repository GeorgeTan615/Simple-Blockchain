package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/blockchain-prac/config"
	"github.com/blockchain-prac/utils"
	"github.com/ethereum/go-ethereum/crypto"
)

var W *Wallet

type Wallet struct {
	Balance      int
	PrivateKey   *ecdsa.PrivateKey
	PublicKeyStr string
}

func NewWallet() *Wallet {
	privateKey := utils.GeneratePrivateKey()
	return &Wallet{
		Balance:      config.INITIAL_BALANCE,
		PrivateKey:   privateKey,
		PublicKeyStr: string(crypto.FromECDSAPub(&privateKey.PublicKey)),
	}
}

func NewBlockchainWallet() *Wallet {
	wallet := NewWallet()
	wallet.Balance = 1000000000
	wallet.PublicKeyStr = "blockchain-wallet"
	return wallet
}

func (w *Wallet) String() string {
	return fmt.Sprintf(
		"Wallet -\nBalance: %d\nPublicKey: %x\n",
		w.Balance,
		w.PublicKeyStr,
	)
}

func (w *Wallet) Sign(data []byte) (string, error) {
	sig, err := crypto.Sign(data, w.PrivateKey)
	if err != nil {
		return "", err
	}
	return string(sig), nil
}

func (w *Wallet) CreateTransaction(recipient string, amount int, transactionPool *TransactionPool) (*Transaction, error) {
	if amount > w.Balance {
		return nil, errors.New("Amount exceeded balance")
	}

	existingTransaction := transactionPool.FindExistingTransactionByPubKey(w.PublicKeyStr)

	if existingTransaction != nil {
		existingTransaction.Update(w, recipient, amount)
		return existingTransaction, nil
	}

	newTransaction, err := NewTransaction(w, recipient, amount)

	if err != nil {
		return nil, errors.New("Error creating new transaction")
	}

	transactionPool.UpsertTransaction(newTransaction)
	return newTransaction, nil
}
