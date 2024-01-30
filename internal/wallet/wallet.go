package wallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/blockchain-prac/utils"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	InitialBalance = 500
)

type Wallet struct {
	Balance      int
	PrivateKey   *ecdsa.PrivateKey
	PublicKeyStr string
}

func NewWallet() *Wallet {
	privateKey := utils.GeneratePrivateKey()
	return &Wallet{
		Balance:      InitialBalance,
		PrivateKey:   privateKey,
		PublicKeyStr: string(crypto.FromECDSAPub(&privateKey.PublicKey)),
	}
}

func (w *Wallet) String() string {
	return fmt.Sprintf(
		"Wallet -\nBalance: %d\nPublicKey: %x\n",
		w.Balance,
		w.PublicKeyStr,
	)
}
