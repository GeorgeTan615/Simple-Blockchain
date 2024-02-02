package blockchain

import "github.com/blockchain-prac/internal/wallet"

type AddBlockRequest struct {
	Data []*wallet.Transaction `json:"data"`
}

type CreateTransactionRequest struct {
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}
