package blockchain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChainTestSuite struct {
	suite.Suite
	bc *Blockchain
	w  *Wallet
	t  *Transaction
}

func (s *ChainTestSuite) SetupTest() {
	s.bc = NewBlockchain()
	s.w = NewWallet()
	s.t, _ = NewTransaction(s.w, "recipient", 50)
}

func (s *ChainTestSuite) TestNewBlockchainAddBlock() {
	s.Equal(1, len(s.bc.Chain))
	s.Equal(NewGenesisBlock(), s.bc.Chain[0])
}

func (s *ChainTestSuite) TestBlockchainAddBlock() {
	data := []*Transaction{s.t}

	s.bc.AddBlock(data)
	s.Equal(data, s.bc.Chain[len(s.bc.Chain)-1].Data)
}

func (s *ChainTestSuite) TestBlockchainValidatesValidChain() {
	bc2 := NewBlockchain()
	data := []*Transaction{s.t}

	bc2.AddBlock(data)
	s.True(s.bc.IsValidChain(bc2.Chain))
}

func (s *ChainTestSuite) TestBlockchainInvalidatesCorruptedGenesisBlock() {
	bc2 := NewBlockchain()

	bc2.Chain[0].Data = append(bc2.Chain[0].Data, s.t)
	s.False(s.bc.IsValidChain(bc2.Chain))
}

func (s *ChainTestSuite) TestBlockchainInvalidatesCorruptedChain() {
	bc2 := NewBlockchain()
	data := []*Transaction{s.t}

	bc2.AddBlock(data)

	newTransaction, newErr := NewTransaction(s.w, "recipient", 100)
	s.Nil(newErr)
	newData := []*Transaction{newTransaction}

	bc2.Chain[1].Data = newData
	s.False(s.bc.IsValidChain(bc2.Chain))
}

func (s *ChainTestSuite) TestBlockchainReplaceValidChain() {
	bc2 := NewBlockchain()
	data := []*Transaction{s.t}

	bc2.AddBlock(data)
	s.bc.ReplaceChain(bc2.Chain)
	s.Equal(bc2.Chain, s.bc.Chain)
}

func (s *ChainTestSuite) TestBlockchainCantReplaceShorterChain() {
	data := []*Transaction{s.t}

	s.bc.AddBlock(data)
	bc2 := NewBlockchain()
	s.bc.ReplaceChain(bc2.Chain)
	s.NotEqual(bc2.Chain, s.bc.Chain)
}

func (s *ChainTestSuite) TestBlockchainCantReplaceInvalidChain() {
	bc2 := NewBlockchain()
	data := []*Transaction{s.t}

	bc2.AddBlock(data)
	bc2.Chain[1].Hash = "123"
	s.bc.ReplaceChain(bc2.Chain)
	s.NotEqual(bc2.Chain, s.bc.Chain)
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ChainTestSuite))
}
