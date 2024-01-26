package blockchain

import (
	"testing"

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
