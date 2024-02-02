package main

import (
	"github.com/blockchain-prac/internal/blockchain"
	"github.com/blockchain-prac/internal/miner"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/blocks", blockchain.GetBlocksController)
	r.POST("/blocks", blockchain.AddBlockController)

	r.GET("/", blockchain.P2PServerInstance.Listen)

	r.GET("/transactions", blockchain.GetTransactionsController)
	r.POST("/transactions", blockchain.CreateTransactionController)

	r.GET("/public-key", blockchain.GetPublicKeyController)
	r.POST("/mine-transactions", miner.MineTransactionsController)
}
