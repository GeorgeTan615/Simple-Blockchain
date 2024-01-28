package blockchain

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMineBlock(t *testing.T) {
	fooBlock := MineBlock(NewGenesisBlock(), []string{"foo"})
	t.Log(fooBlock)
}

func TestBlockOps(t *testing.T) {
	data := []string{"foo", "bar"}
	lastBlock := NewGenesisBlock()
	nextBlock := MineBlock(lastBlock, data)

	assert.Equal(t, data, nextBlock.Data)
	assert.Equal(t, lastBlock.Hash, nextBlock.LastHash)
}

func TestProofOfWork(t *testing.T) {
	testInput := []struct {
		difficulty int
		data       []string
		lastHash   string
	}{
		{
			difficulty: 1,
			data:       []string{"1", "2"},
			lastHash:   "123",
		},
		{
			difficulty: 3,
			data:       []string{"3", "4"},
			lastHash:   "456",
		},
		{
			difficulty: 5,
			data:       []string{"hello", "world"},
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
		assert.Equal(t, strings.Repeat("0", input.difficulty), resp.hash[:input.difficulty])
	}
}

func TestAdjustDifficulty(t *testing.T) {
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
		assert.Equal(t, input.expectedDifficulty, finalDifficulty)
	}
}
