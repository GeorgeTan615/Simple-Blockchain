package api

import (
	"github.com/blockchain-prac/blockchain"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/blocks", blockchain.GetBlocksController)
	r.POST("/blocks", blockchain.AddBlockController)
}
