package blockchain

import (
	"net/http"

	"github.com/blockchain-prac/pkg"
	"github.com/gin-gonic/gin"
)

func GetBlocksController(c *gin.Context) {
	c.JSON(200, Bc.Chain)
}

func AddBlockController(c *gin.Context) {
	var req AddBlockRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResp("Bad Request"))
	}

	Bc.AddBlock(req.Data)
	c.Redirect(http.StatusFound, "/blocks")
}
