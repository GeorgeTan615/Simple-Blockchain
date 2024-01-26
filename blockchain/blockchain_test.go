package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockchainAddBlock(t *testing.T) {
	bc := NewBlockchain()
	assert.Equal(t, 1, len(bc.Chain))
	assert.Equal(t, NewGenesisBlock(), bc.Chain[0])
}

func TestBlockchainAddBlock(t *testing.T) {
	bc := NewBlockchain()
	data := []string{"someData"}
	bc.AddBlock(data)
	assert.Equal(t, data, bc.Chain[len(bc.Chain)-1].Data)
}

func TestBlockchainValidatesValidChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()
	bc2.AddBlock([]string{"foo"})
	assert.True(t, bc.IsValidChain(bc2.Chain))
}

func TestBlockchainInvalidatesCorruptedGenesisBlock(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()
	bc2.Chain[0].Data = append(bc2.Chain[0].Data, "123")
	assert.False(t, bc.IsValidChain(bc2.Chain))
}

func TestBlockchainInvalidatesCorruptedChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()
	bc2.AddBlock([]string{"bar"})
	bc2.Chain[1].Data = []string{"baz"}
	assert.False(t, bc.IsValidChain(bc2.Chain))
}

func TestBlockchanReplaceValidChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()
	bc2.AddBlock([]string{"123", "456"})
	bc.ReplaceChain(bc2.Chain)
	assert.Equal(t, bc2.Chain, bc.Chain)
}

func TestBlockchanCantReplaceShorterChain(t *testing.T) {
	bc := NewBlockchain()
	bc.AddBlock([]string{"123", "456"})
	bc2 := NewBlockchain()
	bc.ReplaceChain(bc2.Chain)
	assert.NotEqual(t, bc2.Chain, bc.Chain)
}

func TestBlockchanCantReplaceInvalidChain(t *testing.T) {
	bc := NewBlockchain()
	bc2 := NewBlockchain()
	bc2.AddBlock([]string{"123", "456"})
	bc2.Chain[1].Hash = "123"
	bc.ReplaceChain(bc2.Chain)
	assert.NotEqual(t, bc2.Chain, bc.Chain)
}
