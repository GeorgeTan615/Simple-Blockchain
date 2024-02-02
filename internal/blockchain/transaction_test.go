package blockchain

import (
	"testing"

	"github.com/blockchain-prac/config"
	"github.com/stretchr/testify/suite"
)

type TransactionTestSuite struct {
	suite.Suite
	wallet           *Wallet
	amount           int
	recipientAddress string
	transaction      *Transaction
}

func (s *TransactionTestSuite) SetupTest() {
	wallet := NewWallet()
	amount := 50
	recipientAddress := "recipient"
	transaction, _ := NewTransaction(wallet, recipientAddress, amount)

	s.wallet = wallet
	s.amount = amount
	s.recipientAddress = recipientAddress
	s.transaction = transaction
}

func (s *TransactionTestSuite) TestAmountDeductedFromWalletBalance() {
	for _, output := range s.transaction.Outputs {
		if output.Address == s.wallet.PublicKeyStr {
			s.Equal(s.wallet.Balance-s.amount, output.Amount)
		}

		if output.Address == s.recipientAddress {
			s.Equal(s.amount, output.Amount)
		}
	}
}

func (s *TransactionTestSuite) TestAmountExceedsWalletBalance() {
	transaction, err := NewTransaction(s.wallet, s.recipientAddress, s.wallet.Balance+10000)
	s.Nil(transaction)
	s.NotNil(err)
}

func (s *TransactionTestSuite) TestInputHasBeenCreatedInTransaction() {
	s.Equal(s.wallet.Balance, s.transaction.Input.Amount)
}

func (s *TransactionTestSuite) TestVerifyTransactionSuccess() {
	s.True(VerifyTransaction(s.transaction))
}

func (s *TransactionTestSuite) TestVerifyTransactionFailed_WhenSenderPublicKeyTampered() {
	s.transaction.Input.Address = "random"
	s.False(VerifyTransaction(s.transaction))
}

func (s *TransactionTestSuite) TestVerifyTransactionFailed_WhenTransactionDataTampered() {
	s.transaction.Outputs[0].Amount = 100000
	s.False(VerifyTransaction(s.transaction))
}

func (s *TransactionTestSuite) TestUpdateTransaction() {
	initialTransactionAmount := s.amount
	newTransactionAmount := 20
	err := s.transaction.Update(s.wallet, "newRecipient", newTransactionAmount)

	s.Nil(err)

	var senderAmount int
	for _, output := range s.transaction.Outputs {
		if output.Address == s.wallet.PublicKeyStr {
			senderAmount = output.Amount
		}
	}

	s.Equal(3, len(s.transaction.Outputs))
	s.Equal(
		s.wallet.Balance-initialTransactionAmount-newTransactionAmount,
		senderAmount)
	s.True(VerifyTransaction(s.transaction))
}

func (s *TransactionTestSuite) TestRewardTransaction() {
	bcWallet := NewBlockchainWallet()
	transaction, err := RewardTransaction(s.wallet, bcWallet)
	s.Nil(err)

	for _, output := range transaction.Outputs {
		if output.Address == s.wallet.PublicKeyStr {
			s.Equal(config.MINING_REWARD, output.Amount)
		}
	}
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}
