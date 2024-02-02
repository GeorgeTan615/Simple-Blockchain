package blockchain

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"time"

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

func (w *Wallet) CreateTransaction(recipient string, amount int, blockchain *Blockchain, transactionPool *TransactionPool) (*Transaction, error) {
	w.Balance = w.CalculateWalletBalance(blockchain)

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

func (w *Wallet) CalculateWalletBalance(blockchain *Blockchain) int {
	// Find the latest transaction this wallet has made
	// The latest transaction would hold the latest balance the wallet has, up until the point of time the transaction was made
	transactions := getTransactionsFromBlockchain(blockchain)
	latestTransactionMadeByWallet := getLatestTransactionMadeByWallet(w, transactions)

	var amount int
	var currTime time.Time
	if latestTransactionMadeByWallet == nil {
		amount = w.Balance
		currTime = time.Time{}
	} else {
		for _, output := range latestTransactionMadeByWallet.Outputs {
			if output.Address == w.PublicKeyStr {
				amount = output.Amount
			}
		}
		currTime = *latestTransactionMadeByWallet.Input.Timestamp
	}

	// To get the received amount, loop through all transactions after the latest point of truth,
	// which will not be made by wallet
	for _, transaction := range transactions {
		if transaction.Input.Timestamp.After(currTime) {
			for _, output := range transaction.Outputs {
				if output.Address == w.PublicKeyStr {
					amount += output.Amount
				}
			}
		}
	}

	return amount
}

func getTransactionsFromBlockchain(blockchain *Blockchain) []*Transaction {
	var transactions []*Transaction

	for _, block := range blockchain.Chain {
		transactions = append(transactions, block.Data...)
	}

	return transactions
}

func getLatestTransactionMadeByWallet(wallet *Wallet, transactions []*Transaction) *Transaction {
	currTime := time.Time{}
	for _, transaction := range transactions {
		if transaction.Input.Address == wallet.PublicKeyStr && transaction.Input.Timestamp.After(currTime) {
			return transaction
		}
	}
	return nil
}
