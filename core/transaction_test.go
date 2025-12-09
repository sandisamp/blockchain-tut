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
	tx.From = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}

// func TestTxEncodeDecode(t *testing.T) {
// 	tx := randomTxWithSignature(t)
// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))
// 	txDecoded := new(Transaction)
// 	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
// 	assert.Equal(t, tx, txDecoded)
// }

func randomTxWithSignature(t *testing.T) *Transaction {
	tx := &Transaction{
		Data: []byte("foo"),
	}
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privKey))
	return tx
}
