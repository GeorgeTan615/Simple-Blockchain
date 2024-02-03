package blockchain

import (
	"net/http"

	"github.com/blockchain-prac/internal/errors"
	"github.com/gin-gonic/gin"
)

// @Summary Get the blocks in the blockchain.
// @Description Get the blocks in the blockchain.
// @Tags Blocks
// @Produce json
// @Success 200 {object} blockchain.Blockchain
// @Router /blocks [get]
func GetBlocksController(c *gin.Context) {
	c.JSON(200, Bc)
}

// @Deprecated
// @Summary Adds a new block in the blockchain.
// @Description Adds a new block in the blockchain.
// @Tags Blocks
// @Accept json
// @Produce json
// @Param   req body  blockchain.AddBlockRequest true "Add Block Request"
// @Success 200 {object} blockchain.Blockchain.Chain
// @Failure 400 {object} errors.ErrorResp
// @Router /blocks [post]
func AddBlockController(c *gin.Context) {
	var req AddBlockRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewErrorResp("Error parsing request"))
	}

	Bc.AddBlock(req.Data)
	P2PServerInstance.SyncChains()

	c.Redirect(http.StatusFound, "/blocks")
}

// @Summary Gets the transactions in the transaction pool.
// @Description Gets the transactions in the transaction pool.
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {object} blockchain.TransactionPool.Transactions
// @Router /transactions [get]
func GetTransactionsController(c *gin.Context) {
	c.JSON(200, struct {
		Transactions []*Transaction `json:"transactions"`
	}{
		Transactions: Tp.Transactions,
	})
}

// @Summary Creates a new transaction and submits it to the transaction pool.
// @Description Creates a new transaction and submits it to the transaction pool.
// @Tags Transactions
// @Accept json
// @Produce json
// @Param   req body  blockchain.CreateTransactionRequest true "Create Transaction Request"
// @Success 200 {object} blockchain.TransactionPool.Transactions
// @Failure 400 {object} errors.ErrorResp
// @Failure 500 {object} errors.ErrorResp
// @Router /transactions [post]
func CreateTransactionController(c *gin.Context) {
	var req CreateTransactionRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewErrorResp("Error parsing request"))
	}

	transaction, err := W.CreateTransaction(req.Recipient, req.Amount, Bc, Tp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewErrorResp("Transaction not created successfully"))
	}

	P2PServerInstance.BroadcastTransaction(transaction)
	c.Redirect(http.StatusFound, "/transactions")
}

// @Summary Gets the public key of the user's wallet.
// @Description Gets the public key of the user's wallet.
// @Tags Wallet
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /public-key [get]
func GetPublicKeyController(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		PublicKey string `json:"publicKey"`
	}{
		PublicKey: W.PublicKeyStr,
	})
}
