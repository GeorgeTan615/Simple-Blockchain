package wallet

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TransactionPoolTestSuite struct {
	suite.Suite
	t  *Transaction
	tp *TransactionPool
	w  *Wallet
}

func (s *TransactionPoolTestSuite) SetupTest() {
	tp := NewTransactionPool()
	wallet := NewWallet()
	amount := 100
	transaction, _ := NewTransaction(wallet, "recipient", amount)
	tp.UpsertTransaction(transaction)

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

func TestTransactionPoolTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionPoolTestSuite))
}
