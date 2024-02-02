package wallet

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WalletTestSuite struct {
	suite.Suite
	w  *Wallet
	tp *TransactionPool
}

func (s *WalletTestSuite) SetupTest() {
	s.w = NewWallet()
	s.tp = NewTransactionPool()
}

func (s *WalletTestSuite) TestSenderBalanceReflectsAfterTwoSameTransactions() {
	recipient := "recipient"
	amount := 50

	_, err := s.w.CreateTransaction(recipient, amount, s.tp)
	s.Nil(err)
	_, err = s.w.CreateTransaction(recipient, amount, s.tp)
	s.Nil(err)

	// Transaction output sender balance should accurately reflect the two transactions
	s.Equal(1, len(s.tp.Transactions))
	s.Equal(3, len(s.tp.Transactions[0].Outputs))
	for _, output := range s.tp.Transactions[0].Outputs {
		if output.Address == s.w.PublicKeyStr {
			s.Equal(s.w.Balance-(amount*2), output.Amount)
		}
	}
}

func (s *WalletTestSuite) TestTwoSameRecipientOutputAfterTwoSameTransactions() {
	recipient := "recipient"
	amount := 50

	_, err := s.w.CreateTransaction(recipient, amount, s.tp)
	s.Nil(err)
	_, err = s.w.CreateTransaction(recipient, amount, s.tp)
	s.Nil(err)

	// Should have two similar recipient transaction outputs
	recipientOutputs := filter[TransactionOutput](s.tp.Transactions[0].Outputs, func(to *TransactionOutput) bool {
		return to.Address == recipient
	})

	s.Equal(2, len(recipientOutputs))
	s.Equal(*recipientOutputs[0], *recipientOutputs[1])
}

func filter[T any](slice []*T, f func(*T) bool) []*T {
	var output []*T
	for _, e := range slice {
		if f(e) {
			output = append(output, e)
		}
	}
	return output
}

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}
