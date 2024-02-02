package blockchain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TransactionPoolTestSuite struct {
	suite.Suite
	bc *Blockchain
	t  *Transaction
	tp *TransactionPool
	w  *Wallet
}

func (s *TransactionPoolTestSuite) SetupTest() {
	bc := NewBlockchain()
	tp := NewTransactionPool()
	wallet := NewWallet()
	amount := 100
	transaction, _ := wallet.CreateTransaction("recipient", amount, bc, tp)

	s.bc = bc
	s.t = transaction
	s.tp = tp
	s.w = wallet
}

func (s *TransactionPoolTestSuite) TestTransactionAddedToPool() {
	s.Equal(1, len(s.tp.Transactions))
	s.Equal(s.t.Id, s.tp.Transactions[0].Id)
	s.Equal(s.t, s.tp.Transactions[0])
}

func (s *TransactionPoolTestSuite) TestTransactionUpdated() {
	oldTransaction := *s.t
	err := s.t.Update(s.w, "newRecipient", 10)
	s.Nil(err)
	s.NotEqual(oldTransaction, *s.t)

	s.tp.UpsertTransaction(s.t)

	for _, t := range s.tp.Transactions {
		if t.Id == s.t.Id {
			s.NotEqual(oldTransaction, *t)
		}
	}
}

func (s *TransactionPoolTestSuite) TestGetValidTransactionsWithNonCorruptedTransactions() {
	expectedValidTransactions := s.tp.Transactions
	for i := 0; i < 6; i++ {
		newWallet := NewWallet()
		transaction, err := newWallet.CreateTransaction("recipient2", 100, s.bc, s.tp)

		s.Nil(err)
		if i%2 == 0 {
			transaction.Outputs[0].Amount = 999999
		} else {
			expectedValidTransactions = append(expectedValidTransactions, transaction)
		}
	}

	currentValidTransactions := s.tp.GetValidTransactions()

	s.NotEqual(len(s.tp.Transactions), len(currentValidTransactions))
	s.True(reflect.DeepEqual(expectedValidTransactions, currentValidTransactions))
}

func (s *TransactionPoolTestSuite) TestClear() {
	s.tp.Clear()
	s.Equal(0, len(s.tp.Transactions))
}

func TestTransactionPoolTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionPoolTestSuite))
}
