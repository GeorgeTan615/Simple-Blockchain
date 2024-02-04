package main

import (
	"github.com/blockchain-prac/cmd/docs"
	"github.com/blockchain-prac/internal/blockchain"
	"github.com/blockchain-prac/internal/miner"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/blocks", blockchain.GetBlocksController)
	r.POST("/blocks", blockchain.AddBlockController)

	r.GET("/", blockchain.P2PServerInstance.Listen)

	r.GET("/transactions", blockchain.GetTransactionsController)
	r.POST("/transactions", blockchain.CreateTransactionController)

	r.GET("/public-key", blockchain.GetPublicKeyController)
	r.POST("/mine-transactions", miner.MineTransactionsController)

	// Swagger
	docs.SwaggerInfo.Title = "Simple Blockchain"
	docs.SwaggerInfo.Description = "This is a project built to gain more understanding on the fundamental workings of blockchain."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
