package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/blockchain-prac/config"
	"github.com/blockchain-prac/utils"
)

type Block struct {
	Timestamp  *time.Time     `json:"timestamp"`
	LastHash   string         `json:"lastHash"`
	Hash       string         `json:"hash"`
	Data       []*Transaction `json:"data"`
	Nonce      int            `json:"nonce"`
	Difficulty int            `json:"difficulty"`
}

func NewBlock(timestamp *time.Time, lastHash, hash string, data []*Transaction, nonce, difficulty int) *Block {
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
	return fmt.Sprintf("Block -\nTimestamp: %s\nLast Hash: %s\nHash: %s\nData: %v\n Nonce: %d\n Difficulty: %d",
		b.Timestamp,
		b.LastHash,
		b.Hash,
		b.Data,
		b.Nonce,
		b.Difficulty)
}

func NewGenesisBlock() *Block {
	currTime := time.Date(2023, time.December, 3, 0, 0, 0, 0, time.UTC)
	return NewBlock(&currTime, "", "f1r57-h45h", nil, 0, config.DIFFICULTY)
}

func MineBlock(lastBlock *Block, data []*Transaction) *Block {
	difficulty := adjustDifficulty(lastBlock, time.Now(), config.MINE_RATE)
	resp := proofOfWork(&ProofOfWorkReq{
		nonce:      0,
		difficulty: difficulty,
		lastHash:   lastBlock.Hash,
		data:       data,
	})

	return NewBlock(&resp.createdAt, lastBlock.Hash, resp.hash, data, resp.nonce, difficulty)
}

func Hash(timestamp *time.Time, lastHash string, data []byte, nonce, difficulty int) string {
	hashInput := fmt.Sprintf("%s%s%s%d%d", timestamp.Format(time.RFC3339), lastHash, hex.EncodeToString(data), nonce, difficulty)
	return fmt.Sprintf("%x", string(utils.Hash([]byte(hashInput))))
}

func BlockHash(block *Block) string {
	dataBytes, err := json.Marshal(block.Data)
	if err != nil {
		return ""
	}
	return Hash(block.Timestamp, block.LastHash, dataBytes, block.Nonce, block.Difficulty)
}
