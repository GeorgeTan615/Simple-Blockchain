package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/blockchain-prac/internal/blockchain"
	"github.com/blockchain-prac/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	blockchain.Bc = blockchain.NewBlockchain()
	blockchain.P2PServerInstance = blockchain.NewP2PServer(blockchain.Bc, []*websocket.Conn{})
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

			blockchain.P2PServerInstance.Sockets = append(blockchain.P2PServerInstance.Sockets, conn)
			blockchain.HandleBlockchainUpdates(conn, u.Host)
		}(peer)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	r := gin.Default()

	InitRoutes(r)

	port := utils.GetEnv("HTTP_PORT", "8080")
	wsPeers := strings.Split(utils.GetEnv("PEERS", ""), ",")

	if peersEnvVar := utils.GetEnv("PEERS", ""); peersEnvVar != "" {
		wsPeers = strings.Split(peersEnvVar, ",")
		connectToWsPeers(wsPeers)
	}

	r.Run(fmt.Sprintf(":%s", port))
}
