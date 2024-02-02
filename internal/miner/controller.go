package miner

import (
	"net/http"

	"github.com/blockchain-prac/internal/errors"
	"github.com/gin-gonic/gin"
)

func MineTransactionsController(c *gin.Context) {
	_, err := M.Mine()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewErrorResp("Mine Block Failure"))
	}
	c.Redirect(http.StatusFound, "/blocks")
}
