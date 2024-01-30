package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

const (
	Difficulty = 4
	MineRate   = 1 * time.Second
)

type Block struct {
	Timestamp  *time.Time
	LastHash   string
	Hash       string
	Data       []string
	Nonce      int
	Difficulty int
}

func NewBlock(timestamp *time.Time, lastHash, hash string, data []string, nonce, difficulty int) *Block {
	return &Block{
		Timestamp:  timestamp,
		LastHash:   lastHash,
		Hash:       hash,
		Data:       data,
		Nonce:      nonce,
		Difficulty: difficulty,
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("Block -\nTimestamp: %s\nLast Hash: %s\nHash: %s\nData: %s\n Nonce: %d\n Difficulty: %d",
		b.Timestamp,
		b.LastHash,
		b.Hash,
		b.Data,
		b.Nonce,
		b.Difficulty)
}

func NewGenesisBlock() *Block {
	currTime := time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC)
	return NewBlock(&currTime, "", "f1r57-h45h", []string{}, 0, Difficulty)
}

func MineBlock(lastBlock *Block, data []string) *Block {
	difficulty := adjustDifficulty(lastBlock, time.Now(), MineRate)
	resp := proofOfWork(&ProofOfWorkReq{
		nonce:      0,
		difficulty: difficulty,
		lastHash:   lastBlock.Hash,
		data:       data,
	})

	return NewBlock(&resp.createdAt, lastBlock.Hash, resp.hash, data, resp.nonce, difficulty)
}

func Hash(timestamp *time.Time, lastHash string, data []string, nonce, difficulty int) string {
	hashInput := fmt.Sprintf("%s%s%s%d%d", timestamp.Format(time.RFC3339), lastHash, data, nonce, difficulty)
	sum := sha256.Sum256([]byte(hashInput))
	return fmt.Sprintf("%x", sum)
}

func BlockHash(block *Block) string {
	timestamp, lastHash, data, nonce := block.Timestamp, block.LastHash, block.Data, block.Nonce
	return Hash(timestamp, lastHash, data, nonce, block.Difficulty)
}