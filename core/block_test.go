package core

import (
	"testing"
	"time"

	"github.com/sandisamp/blockchain-tut/crypto"
	"github.com/sandisamp/blockchain-tut/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:        1,
		PrevbBlockHash: types.RandomHash(),
		Timestamp:      time.Now().UnixNano(),
		Height:         height,
	}
	tx := Transaction{
		Data: []byte("foo"),
	}
	return NewBlock(header, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {
	b := randomBlock(0)

	privKey := crypto.GeneratePrivateKey()
	if err := b.Sign(privKey); err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(0)
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
