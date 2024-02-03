package blockchain

import (
	"github.com/blockchain-prac/utils"
)

var Tp *TransactionPool

type TransactionPool struct {
	Transactions []*Transaction
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{
		Transactions: nil,
	}
}

func (tp *TransactionPool) UpsertTransaction(transaction *Transaction) {
	idx := findTransactionIdxById(tp.Transactions, transaction.Id)

	if idx != -1 {
		tp.Transactions[idx] = transaction
		return
	}

	tp.Transactions = append(tp.Transactions, transaction)
}

func findTransactionIdxById(transactions []*Transaction, id string) int {
	for i, transaction := range transactions {
		if transaction.Id == id {
			return i
		}
	}

	return -1
}

func (tp *TransactionPool) FindExistingTransactionByPubKey(senderAddress string) *Transaction {
	for _, transaction := range tp.Transactions {
		if transaction.Input.Address == senderAddress {
			return transaction
		}
	}

	return nil
}

func (tp *TransactionPool) GetValidTransactions() []*Transaction {
	return utils.Filter[Transaction](tp.Transactions, isValidTransaction)
}

func isValidTransaction(transaction *Transaction) bool {
	// Verify if total amount tallies
	totalAmount := 0
	for _, output := range transaction.Outputs {
		totalAmount += output.Amount
	}
	if totalAmount != transaction.Input.Amount {
		return false
	}

	// Verify transaction signature
	return VerifyTransaction(transaction)
}

func (tp *TransactionPool) Clear() {
	tp.Transactions = nil
}
