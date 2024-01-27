package main

import (
	"github.com/blockchain-prac/internal/blockchain"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/blocks", blockchain.GetBlocksController)
	r.POST("/blocks", blockchain.AddBlockController)
	r.GET("/", blockchain.P2PServerInstance.Listen)
}
