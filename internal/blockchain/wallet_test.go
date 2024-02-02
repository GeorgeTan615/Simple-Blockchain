package blockchain

import (
	"testing"

	"github.com/blockchain-prac/utils"
	"github.com/stretchr/testify/suite"
)

type WalletTestSuite struct {
	suite.Suite
	bc *Blockchain
	w  *Wallet
	tp *TransactionPool
}

func (s *WalletTestSuite) SetupTest() {
	s.bc = NewBlockchain()
	s.w = NewWallet()
	s.tp = NewTransactionPool()
}

func (s *WalletTestSuite) TestSenderBalanceReflectsAfterTwoSameTransactions() {
	recipient := "recipient"
	amount := 50

	_, err := s.w.CreateTransaction(recipient, amount, s.bc, s.tp)
	s.Nil(err)
	_, err = s.w.CreateTransaction(recipient, amount, s.bc, s.tp)
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

	_, err := s.w.CreateTransaction(recipient, amount, s.bc, s.tp)
	s.Nil(err)
	_, err = s.w.CreateTransaction(recipient, amount, s.bc, s.tp)
	s.Nil(err)

	recipientOutputs := utils.Filter[TransactionOutput](s.tp.Transactions[0].Outputs, func(to *TransactionOutput) bool {
		return to.Address == recipient
	})

	s.Equal(2, len(recipientOutputs))
	s.Equal(*recipientOutputs[0], *recipientOutputs[1])
}

func (s *WalletTestSuite) TestWalletCalculateBalance() {
	currWalletInitialBalance := s.w.Balance

	anotherWallet := NewWallet()
	anotherWalletInitialBalance := anotherWallet.Balance
	amount1 := 50

	_, err := anotherWallet.CreateTransaction(s.w.PublicKeyStr, amount1, s.bc, s.tp)
	s.Nil(err)

	s.bc.AddBlock(s.tp.Transactions)
	s.Equal(anotherWallet.Balance-amount1, anotherWallet.CalculateWalletBalance(s.bc))
	s.Equal(s.w.Balance+amount1, s.w.CalculateWalletBalance(s.bc))

	s.tp.Clear()

	amount2 := 100
	_, err = s.w.CreateTransaction(anotherWallet.PublicKeyStr, amount2, s.bc, s.tp)
	s.Nil(err)

	s.bc.AddBlock(s.tp.Transactions)
	s.Equal(anotherWalletInitialBalance-amount1+amount2, anotherWallet.CalculateWalletBalance(s.bc))
	s.Equal(currWalletInitialBalance+amount1-amount2, s.w.CalculateWalletBalance(s.bc))
}

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}
