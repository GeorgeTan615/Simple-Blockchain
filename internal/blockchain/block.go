package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Timestamp *time.Time
	LastHash  string
	Hash      string
	Data      []string
}

func (b *Block) String() string {
	return fmt.Sprintf("Block -\nTimestamp: %s\nLast Hash: %s\nHash: %s\nData: %s",
		b.Timestamp,
		b.LastHash,
		b.Hash,
		b.Data)
}

func NewBlock(timestamp *time.Time, lastHash, hash string, data []string) *Block {
	return &Block{
		Timestamp: timestamp,
		LastHash:  lastHash,
		Hash:      hash,
		Data:      data,
	}
}

func NewGenesisBlock() *Block {
	currTime := time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC)
	return NewBlock(&currTime, "", "f1r57-h45h", []string{})
}

func MineBlock(lastBlock *Block, data []string) *Block {
	currTime := time.Now()
	lastHash := lastBlock.Hash
	return NewBlock(&currTime, lastHash, Hash(&currTime, lastHash, data), data)
}

func Hash(timestamp *time.Time, lastHash string, data []string) string {
	hashInput := fmt.Sprintf("%s%s%s", timestamp.Format(time.RFC3339), lastHash, data)
	sum := sha256.Sum256([]byte(hashInput))
	return fmt.Sprintf("%x", sum)
}

func BlockHash(block *Block) string {
	timestamp, lastHash, data := block.Timestamp, block.LastHash, block.Data
	return Hash(timestamp, lastHash, data)
}
