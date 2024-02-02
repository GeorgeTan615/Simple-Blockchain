package miner

import (
	"github.com/blockchain-prac/internal/blockchain"
)

var M *Miner

type Miner struct {
	Blockchain      *blockchain.Blockchain
	TransactionPool *blockchain.TransactionPool
	Wallet          *blockchain.Wallet
	P2PServer       *blockchain.P2PServer
}

func NewMiner(
	blockchain *blockchain.Blockchain,
	transactionPool *blockchain.TransactionPool,
	wallet *blockchain.Wallet,
	p2pserver *blockchain.P2PServer,
) *Miner {
	return &Miner{
		Blockchain:      blockchain,
		TransactionPool: transactionPool,
		Wallet:          wallet,
		P2PServer:       p2pserver,
	}
}

func (m *Miner) Mine() (*blockchain.Block, error) {
	validTransactions := m.TransactionPool.GetValidTransactions()
	rewardTransaction, err := blockchain.RewardTransaction(m.Wallet, blockchain.NewBlockchainWallet())

	if err != nil {
		return nil, err
	}

	validTransactions = append(validTransactions, rewardTransaction)
	block := m.Blockchain.AddBlock(validTransactions)
	m.P2PServer.SyncChains()
	m.TransactionPool.Clear()
	m.P2PServer.BroadcastClearTransactions()

	return block, nil
}
