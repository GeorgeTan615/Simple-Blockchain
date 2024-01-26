package blockchain

import "reflect"

var Bc *Blockchain

type Blockchain struct {
	Chain []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain: []*Block{
			NewGenesisBlock(),
		},
	}
}

func (bc *Blockchain) AddBlock(data []string) *Block {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := MineBlock(lastBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
	return newBlock
}

func (bc *Blockchain) IsValidChain(chain []*Block) bool {
	if len(chain) < 1 || !reflect.DeepEqual(chain[0], NewGenesisBlock()) {
		return false
	}

	for i := 1; i < len(chain); i++ {
		currBlock := chain[i]
		prevBlock := chain[i-1]

		if currBlock.LastHash != prevBlock.Hash || currBlock.Hash != BlockHash(currBlock) {
			return false
		}
	}

	return true
}

func (bc *Blockchain) ReplaceChain(newChain []*Block) {
	if len(newChain) > len(bc.Chain) && bc.IsValidChain(newChain) {
		bc.Chain = newChain
	}
}
