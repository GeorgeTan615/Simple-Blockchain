package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/blockchain-prac/utils"
	"github.com/ethereum/go-ethereum/crypto"
)

type TransactionInput struct {
	Timestamp *time.Time
	Amount    int
	Address   string
	Signature string
}

type TransactionOutput struct {
	Amount  int
	Address string
}

type Transaction struct {
	Id      string
	Input   *TransactionInput
	Outputs []*TransactionOutput
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

	transaction := &Transaction{
		Id:      utils.GenerateUniqueId(),
		Input:   nil,
		Outputs: outputs,
	}

	err := signTransaction(senderWallet, transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
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
