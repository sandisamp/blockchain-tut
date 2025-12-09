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
	// firstSeen is the timestamp of when tx is first seen locally
	firstSeen int64
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

func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}

func (tx *Transaction) SetFirstSeen(firstSeen int64) {
	tx.firstSeen = firstSeen
}

func (tx *Transaction) FirstSeen() int64 {
	return tx.firstSeen
}
