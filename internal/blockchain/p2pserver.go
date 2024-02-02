package blockchain

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type P2PServerMessage struct {
	Blockchain           *Blockchain
	Transaction          *Transaction
	ClearTransactionPool bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var P2PServerInstance *P2PServer

type P2PServer struct {
	Blockchain      *Blockchain
	TransactionPool *TransactionPool
	Sockets         []*websocket.Conn
}

func NewP2PServer(blockchain *Blockchain, transactionPool *TransactionPool, sockets []*websocket.Conn) *P2PServer {
	return &P2PServer{
		Blockchain:      blockchain,
		TransactionPool: transactionPool,
		Sockets:         sockets,
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

	msg := &P2PServerMessage{
		Blockchain: Bc,
	}

	if err := conn.WriteJSON(msg); err != nil {
		log.Println("Error sending blockchain:", err)
		return
	}
	log.Println("Current blockchain sent")

	HandleBlockchainUpdates(conn, clientAddress)
}

func (p2pServer *P2PServer) SyncChains() {
	msg := &P2PServerMessage{
		Blockchain: Bc,
	}

	for _, wsConn := range p2pServer.Sockets {
		if err := wsConn.WriteJSON(msg); err != nil {
			log.Println("Error sending blockchain:", wsConn, err)
			return
		}
		log.Println("Chains synced")
	}
}

func (p2pServer *P2PServer) BroadcastTransaction(transaction *Transaction) {
	msg := &P2PServerMessage{
		Transaction: transaction,
	}

	for _, wsConn := range p2pServer.Sockets {
		if err := wsConn.WriteJSON(msg); err != nil {
			log.Println("Error broadcasting transaction:", wsConn, err)
			return
		}
		log.Println("Transaction broadcasted")
	}
}

func (p2pServer *P2PServer) BroadcastClearTransactions() {
	msg := &P2PServerMessage{
		ClearTransactionPool: true,
	}

	for _, wsConn := range p2pServer.Sockets {
		if err := wsConn.WriteJSON(msg); err != nil {
			log.Println("Error broadcasting clear transaction:", wsConn, err)
			return
		}
		log.Println("Clear transaction broadcasted")
	}
}

func HandleBlockchainUpdates(conn *websocket.Conn, clientAddress string) {
	for {
		var msg P2PServerMessage
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Error receiving message from:", clientAddress, err)
			return
		}

		if msg.Blockchain != nil {
			ok := P2PServerInstance.Blockchain.ReplaceChain(msg.Blockchain.Chain)
			if ok {
				P2PServerInstance.SyncChains()
			}
		}

		if msg.Transaction != nil {
			P2PServerInstance.TransactionPool.UpsertTransaction(msg.Transaction)
			log.Println("Transaction broadcast received")
		}

		if msg.ClearTransactionPool {
			P2PServerInstance.TransactionPool.Transactions = nil
		}
	}
}
