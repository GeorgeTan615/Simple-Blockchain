package blockchain

type AddBlockRequest struct {
	Data []string `json:"data"`
}

type CreateTransactionRequest struct {
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}
