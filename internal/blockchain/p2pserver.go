package blockchain

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var P2PServerInstance *P2PServer

type P2PServer struct {
	Blockchain *Blockchain
	Sockets    []*websocket.Conn
}

func NewP2PServer(blockchain *Blockchain, sockets []*websocket.Conn) *P2PServer {
	return &P2PServer{
		Blockchain: blockchain,
		Sockets:    sockets,
	}
}

func (p2pServer *P2PServer) Listen(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Websocket upgrade failed", err)
		return
	}

	defer conn.Close()
	clientAddress := c.Request.RemoteAddr
	p2pServer.Sockets = append(p2pServer.Sockets, conn)
	log.Println("Socket connected:", clientAddress)

	if err := conn.WriteJSON(Bc); err != nil {
		log.Println("Error sending blockchain:", err)
		return
	}
	log.Println("Current blockchain sent")

	HandleBlockchainUpdates(conn, clientAddress)
}

func (p2pServer *P2PServer) SyncChains() {
	for _, wsConn := range p2pServer.Sockets {
		if err := wsConn.WriteJSON(Bc); err != nil {
			log.Println("Error sending blockchain:", wsConn, err)
			return
		}
		log.Println("Chains synced")
	}
}

func HandleBlockchainUpdates(conn *websocket.Conn, clientAddress string) {
	for {
		var blockchain Blockchain
		if err := conn.ReadJSON(&blockchain); err != nil {
			log.Println("Error receiving message from:", clientAddress, err)
			return
		}

		ok := P2PServerInstance.Blockchain.ReplaceChain(blockchain.Chain)
		if ok {
			for _, wsConn := range P2PServerInstance.Sockets {
				if err := wsConn.WriteJSON(Bc); err != nil {
					log.Println("Error sending blockchain:", wsConn, err)
					return
				}
				log.Println("Chains synced")
			}
		}
	}

}
