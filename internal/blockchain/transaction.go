package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/blockchain-prac/config"
	"github.com/blockchain-prac/utils"
	"github.com/ethereum/go-ethereum/crypto"
)

type TransactionInput struct {
	Timestamp *time.Time `json:"timestamp"`
	Amount    int        `json:"amount"`
	Address   string     `json:"address"`
	Signature string     `json:"signature"`
}

type TransactionOutput struct {
	Amount  int    `json:"amount"`
	Address string `json:"address"`
}

type Transaction struct {
	Id      string               `json:"id"`
	Input   *TransactionInput    `json:"input"`
	Outputs []*TransactionOutput `json:"outputs"`
}

func NewTransactionWithOutputs(senderWallet *Wallet, outputs []*TransactionOutput) (*Transaction, error) {
	transaction := &Transaction{
		Id:      utils.GenerateUniqueId(),
		Outputs: outputs,
	}
	err := signTransaction(senderWallet, transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func NewTransaction(senderWallet *Wallet, recipientAddress string, amount int) (*Transaction, error) {
	newBalance := senderWallet.Balance - amount
	if newBalance < 0 {
		return nil, errors.New(fmt.Sprintf("Amount %d exceeds wallet balance", senderWallet.Balance))
	}

	outputs := []*TransactionOutput{
		{
			Amount:  newBalance,
			Address: senderWallet.PublicKeyStr,
		},
		{
			Amount:  amount,
			Address: recipientAddress,
		},
	}

	return NewTransactionWithOutputs(senderWallet, outputs)
}

func (t *Transaction) Update(senderWallet *Wallet, recipient string, amount int) error {
	// Loop through outputs to see if amount exceeds sender balance, if not, update accordingly
	for _, output := range t.Outputs {
		if output.Address == senderWallet.PublicKeyStr {
			if amount > output.Amount {
				return errors.New(fmt.Sprintf("Amount %d exceeds wallet balance", amount))
			} else {
				output.Amount -= amount
				break
			}
		}
	}

	// Add into outputs
	t.Outputs = append(t.Outputs, &TransactionOutput{
		Amount:  amount,
		Address: recipient,
	})

	// Re-sign
	err := signTransaction(senderWallet, t)
	return err
}

func signTransaction(senderWallet *Wallet, transaction *Transaction) error {
	outputsBytes, err := json.Marshal(transaction.Outputs)

	if err != nil {
		return errors.New("Error converting transaction outputs to string")
	}

	sig, err := senderWallet.Sign(utils.Hash(outputsBytes))

	if err != nil {
		return errors.New("Sign transaction failed")
	}

	currTime := time.Now()
	transaction.Input = &TransactionInput{
		Timestamp: &currTime,
		Amount:    senderWallet.Balance,
		Address:   senderWallet.PublicKeyStr,
		Signature: sig,
	}

	return nil
}

func VerifyTransaction(transaction *Transaction) bool {
	pubKey := []byte(transaction.Input.Address)

	outputsBytes, err := json.Marshal(transaction.Outputs)

	if err != nil {
		return false
	}

	digestHash := utils.Hash(outputsBytes)
	sigWithRecoveryId := []byte(transaction.Input.Signature)
	sigWithoutRecoveryId := sigWithRecoveryId[:len(sigWithRecoveryId)-1] // remove recovery ID

	return crypto.VerifySignature(pubKey, digestHash, sigWithoutRecoveryId)
}

func RewardTransaction(minerWallet *Wallet, blockchainWallet *Wallet) (*Transaction, error) {
	outputs := []*TransactionOutput{
		{
			Amount:  config.MINING_REWARD,
			Address: minerWallet.PublicKeyStr,
		},
	}
	return NewTransactionWithOutputs(blockchainWallet, outputs)
}
