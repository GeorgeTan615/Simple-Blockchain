package blockchain

type AddBlockRequest struct {
	Data []*Transaction `json:"data"`
}

type CreateTransactionRequest struct {
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}
