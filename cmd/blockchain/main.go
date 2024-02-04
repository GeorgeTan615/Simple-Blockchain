package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/blockchain-prac/internal/blockchain"
	"github.com/blockchain-prac/internal/miner"
	"github.com/blockchain-prac/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func init() {
	blockchain.Bc = blockchain.NewBlockchain()
	blockchain.Tp = blockchain.NewTransactionPool()
	blockchain.W = blockchain.NewWallet()
	blockchain.P2PServerInstance = blockchain.NewP2PServer(blockchain.Bc, blockchain.Tp, nil)
	miner.M = miner.NewMiner(blockchain.Bc, blockchain.Tp, blockchain.W, blockchain.P2PServerInstance)
}

func connectToWsPeers(peers []string) {
	for _, peer := range peers {
		go func(peer string) {
			u := url.URL{
				Scheme: "ws",
				Host:   peer,
				Path:   "/",
			}

			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

			if err != nil {
				log.Fatal("dial:", err)
			}
			log.Println("Connected to:", u.Host)
			defer conn.Close()

			// To enable bidirectional connection to all connected servers
			blockchain.P2PServerInstance.Sockets = append(blockchain.P2PServerInstance.Sockets, conn)
			blockchain.HandleBlockchainUpdates(conn, u.Host)
		}(peer)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	r := gin.Default()
	r.Use(cors.New(cors.DefaultConfig()))

	InitRoutes(r)

	port := utils.GetEnv("HTTP_PORT", "8080")
	wsPeers := strings.Split(utils.GetEnv("PEERS", ""), ",")

	if peersEnvVar := utils.GetEnv("PEERS", ""); peersEnvVar != "" {
		wsPeers = strings.Split(peersEnvVar, ",")
		connectToWsPeers(wsPeers)
	}

	r.Run(fmt.Sprintf(":%s", port))
}
