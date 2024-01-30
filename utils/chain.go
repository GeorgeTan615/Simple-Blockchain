package utils

import (
	"crypto/ecdsa"
	"crypto/sha256"
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

func Hash(hashInput []byte) []byte {
	hash := sha256.Sum256(hashInput)
	return hash[:]
}
