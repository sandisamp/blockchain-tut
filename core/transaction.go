package core

import (
	"fmt"

	"github.com/sandisamp/blockchain-tut/crypto"
	"github.com/sandisamp/blockchain-tut/types"
)

type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature

	// cached version of tx data hash
	hash types.Hash
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) Hash(h Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = h.Hash(tx)
	}

	return tx.hash
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.From = privKey.PublicKey()
	tx.Signature = sig
	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("Transaction has no signature")
	}
	if tx.Signature.Verify(tx.From, tx.Data) {
		return nil
	}
	return fmt.Errorf("Transaction signature is invalid")
}
