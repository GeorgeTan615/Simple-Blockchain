package blockchain

import (
	"net/http"

	"github.com/blockchain-prac/internal/errors"
	"github.com/blockchain-prac/internal/wallet"
	"github.com/gin-gonic/gin"
)

func GetBlocksController(c *gin.Context) {
	c.JSON(200, Bc.Chain)
}

func AddBlockController(c *gin.Context) {
	var req AddBlockRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewErrorResp("Error parsing request"))
	}

	Bc.AddBlock(req.Data)
	P2PServerInstance.SyncChains()

	c.Redirect(http.StatusFound, "/blocks")
}

func GetTransactionsController(c *gin.Context) {
	c.JSON(200, wallet.Tp.Transactions)
}

func CreateTransactionController(c *gin.Context) {
	var req CreateTransactionRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewErrorResp("Error parsing request"))
	}

	transaction, err := wallet.W.CreateTransaction(req.Recipient, req.Amount, wallet.Tp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewErrorResp("Transaction not created successfully"))
	}

	P2PServerInstance.BroadcastTransaction(transaction)
	c.Redirect(http.StatusFound, "/transactions")
}

func GetPublicKeyController(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		PublicKey string `json:"publicKey"`
	}{
		PublicKey: wallet.W.PublicKeyStr,
	})
}
