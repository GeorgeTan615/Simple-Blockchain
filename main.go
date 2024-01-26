package main

import (
	"fmt"
	"log"

	"github.com/blockchain-prac/api"
	"github.com/blockchain-prac/blockchain"
	"github.com/blockchain-prac/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	blockchain.Bc = blockchain.NewBlockchain()
}

func main() {

	r := gin.Default()
	api.InitRoutes(r)
	port := utils.GetEnv("HTTP_PORT", "8080")
	r.Run(fmt.Sprintf(":%s", port))
}
