package wallet

var Tp *TransactionPool

type TransactionPool struct {
	Transactions []*Transaction
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{
		Transactions: []*Transaction{},
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
