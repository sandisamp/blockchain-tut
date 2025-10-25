package core

import (
	"testing"

	"github.com/sandisamp/blockchain-tut/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	tx := Transaction{
		Data: []byte("foo"),
	}
	privKey := crypto.GeneratePrivateKey()
	if err := tx.Sign(privKey); err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifySignature(t *testing.T) {
	tx := Transaction{
		Data: []byte("foo"),
	}
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privKey))

	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}
