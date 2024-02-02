package miner

import (
	"github.com/blockchain-prac/internal/blockchain"
	"github.com/blockchain-prac/internal/wallet"
)

type Miner struct {
	Blockchain      *blockchain.Blockchain
	TransactionPool *wallet.TransactionPool
	Wallet          *wallet.Wallet
	P2PServer       *blockchain.P2PServer
}

func NewMiner(
	blockchain *blockchain.Blockchain,
	transactionPool *wallet.TransactionPool,
	wallet *wallet.Wallet,
	p2pserver *blockchain.P2PServer,
) *Miner {
	return &Miner{
		Blockchain:      blockchain,
		TransactionPool: transactionPool,
		Wallet:          wallet,
		P2PServer:       p2pserver,
	}
}

func (m *Miner) Mine() error {
	validTransactions := m.TransactionPool.GetValidTransactions()
	rewardTransaction, err := wallet.RewardTransaction(m.Wallet, wallet.NewBlockchainWallet())

	if err != nil {
		return err
	}

	validTransactions = append(validTransactions, rewardTransaction)
	m.Blockchain.AddBlock(validTransactions)
	m.P2PServer.SyncChains()
	m.TransactionPool.Clear()

	// Include reward for miner
	// Create a block consisting of valid transactions
	// Synchronize the chains in the peer-to-peer server
	// Clear the transaction pool
	// Broadcast to every miner to clear their transaction

	return nil
}
