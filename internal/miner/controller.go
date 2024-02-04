package miner

import (
	"net/http"

	"github.com/blockchain-prac/internal/errors"
	"github.com/gin-gonic/gin"
)

// @Summary Mine a new block in the blockchain by including the transactions from the transaction pool.
// @Description Mine a new block in the blockchain by including the transactions from the transaction pool.
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {object} blockchain.Blockchain
// @Failure 500 {object} errors.ErrorResp
// @Router /mine-transactions [post]
func MineTransactionsController(c *gin.Context) {
	_, err := M.Mine()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewErrorResp("Mine Block Failure"))
	}
	c.Redirect(http.StatusFound, "/blocks")
}
