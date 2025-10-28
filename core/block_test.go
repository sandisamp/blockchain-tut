package core

import (
	"testing"
	"time"

	"github.com/sandisamp/blockchain-tut/crypto"
	"github.com/sandisamp/blockchain-tut/types"
	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {
	b := randomBlock(0, types.Hash{})

	privKey := crypto.GeneratePrivateKey()
	if err := b.Sign(privKey); err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(0, types.Hash{})
	privKey := crypto.GeneratePrivateKey()
	if err := b.Sign(privKey); err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version:        1,
		PrevbBlockHash: prevBlockHash,
		Timestamp:      time.Now().UnixNano(),
		Height:         height,
	}

	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	b := randomBlock(height, prevBlockHash)
	privKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t, b.Sign(privKey))
	return b
}
