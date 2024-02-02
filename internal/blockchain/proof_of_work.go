package blockchain

import (
	"encoding/json"
	"strings"
	"time"
)

type ProofOfWorkReq struct {
	nonce      int
	difficulty int
	lastHash   string
	data       []*Transaction
}

type ProofOfWorkResp struct {
	nonce     int
	hash      string
	createdAt time.Time
}

func proofOfWork(req *ProofOfWorkReq) *ProofOfWorkResp {
	dataBytes, err := json.Marshal(req.data)

	if err != nil {
		return nil
	}

	for {
		req.nonce++
		currTime := time.Now()
		hash := Hash(&currTime, req.lastHash, dataBytes, req.nonce, req.difficulty)
		if hash[:req.difficulty] == strings.Repeat("0", req.difficulty) {
			return &ProofOfWorkResp{
				nonce:     req.nonce,
				hash:      hash,
				createdAt: currTime,
			}
		}
	}
}

func max(int1, int2 int) int {
	if int1 <= int2 {
		return int2
	} else {
		return int1
	}
}

func adjustDifficulty(lastBlock *Block, currTime time.Time, mineRate time.Duration) int {
	currDifficulty := lastBlock.Difficulty

	if lastBlock.Timestamp.Add(mineRate).After(currTime) {
		currDifficulty++
	} else {
		currDifficulty--
	}

	return max(0, currDifficulty)
}
