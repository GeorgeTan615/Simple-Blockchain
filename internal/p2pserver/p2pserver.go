package p2pserver

import (
	"fmt"
	"log"

	"github.com/blockchain-prac/internal/blockchain"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type P2PServer struct {
	Blockchain *blockchain.Blockchain
	Sockets    []string
}

func NewP2PServer(blockchain *blockchain.Blockchain, sockets []string) *P2PServer {
	return &P2PServer{
		Blockchain: blockchain,
		Sockets:    sockets,
	}
}

func (p2pserver *P2PServer) Listen(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatalln("Websocket upgrade failed", err)
	}

	defer conn.Close()
	clientAddress := c.Request.Host
	p2pserver.Sockets = append(p2pserver.Sockets, clientAddress)
	fmt.Printf("Socket %s connected", clientAddress)
}
