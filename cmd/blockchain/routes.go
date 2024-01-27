package main

import (
	"github.com/blockchain-prac/internal/blockchain"
	"github.com/blockchain-prac/internal/p2pserver"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/blocks", blockchain.GetBlocksController)
	r.POST("/blocks", blockchain.AddBlockController)

	p2pServer := p2pserver.NewP2PServer(blockchain.Bc, []string{})
	r.GET("/", p2pServer.Listen)
}
