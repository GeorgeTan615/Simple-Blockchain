package utils

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func GeneratePrivateKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("Error generating private key", err)
	}
	return key
}

func GenerateUniqueId() string {
	return uuid.New().String()
}
