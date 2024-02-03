package blockchain

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type BlockTestSuite struct {
	suite.Suite
	w *Wallet
	t *Transaction
}

func (s *BlockTestSuite) SetupTest() {
	wallet := NewWallet()
	transaction, _ := NewTransaction(wallet, "recipient", 50)
	s.w = wallet
	s.t = transaction
}

func (s *BlockTestSuite) TestMineBlock() {
	fooBlock := MineBlock(NewGenesisBlock(), []*Transaction{s.t})
	s.T().Log(fooBlock)
}

func (s *BlockTestSuite) TestBlockOps() {
	data := []*Transaction{s.t}
	lastBlock := NewGenesisBlock()
	nextBlock := MineBlock(lastBlock, data)

	s.Equal(data, nextBlock.Data)
	s.Equal(lastBlock.Hash, nextBlock.LastHash)
}

func (s *BlockTestSuite) TestProofOfWork() {
	transaction1 := s.t
	transaction2, _ := NewTransaction(s.w, "recipient2", 10)
	transaction3, _ := NewTransaction(s.w, "recipient3", 20)
	testInput := []struct {
		difficulty int
		data       []*Transaction
		lastHash   string
	}{
		{
			difficulty: 1,
			data:       []*Transaction{transaction1},
			lastHash:   "123",
		},
		{
			difficulty: 3,
			data:       []*Transaction{transaction2},
			lastHash:   "456",
		},
		{
			difficulty: 5,
			data:       []*Transaction{transaction3},
			lastHash:   "abcdef",
		},
	}

	for _, input := range testInput {
		req := &ProofOfWorkReq{
			nonce:      0,
			difficulty: input.difficulty,
			lastHash:   input.lastHash,
			data:       input.data,
		}
		resp := proofOfWork(req)
		s.Equal(strings.Repeat("0", input.difficulty), resp.hash[:input.difficulty])
	}
}

func (s *BlockTestSuite) TestAdjustDifficulty() {
	startTime := time.Date(2024, time.January, 3, 0, 0, 1, 0, time.UTC)
	testInput := []struct {
		lastBlock          *Block
		currTime           time.Time
		mineRate           time.Duration
		expectedDifficulty int
	}{
		// Mined within mine rate, reduce difficulty
		{
			lastBlock: &Block{
				Difficulty: 10,
				Timestamp:  &startTime,
			},
			currTime:           time.Date(2024, time.January, 3, 0, 0, 1, 30, time.UTC),
			mineRate:           1 * time.Second,
			expectedDifficulty: 11,
		},

		// Mined exceeding mine rate, increase difficulty
		{
			lastBlock: &Block{
				Difficulty: 10,
				Timestamp:  &startTime,
			},
			currTime:           time.Date(2024, time.January, 3, 0, 0, 10, 0, time.UTC),
			mineRate:           1 * time.Second,
			expectedDifficulty: 9,
		},
	}
	for _, input := range testInput {
		finalDifficulty := adjustDifficulty(input.lastBlock, input.currTime, input.mineRate)
		s.Equal(input.expectedDifficulty, finalDifficulty)
	}
}

func TestBlockTestSuite(t *testing.T) {
	suite.Run(t, new(BlockTestSuite))
}
