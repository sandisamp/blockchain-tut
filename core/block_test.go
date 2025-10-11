package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/sandisamp/blockchain-tut/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:    1,
		PrevbBlock: types.RandomHash(),
		Timestamp:  time.Now().UnixNano(),
		Height:     10,
		Nonce:      99234,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h.Version, hDecode.Version)
	assert.Equal(t, h.PrevbBlock, hDecode.PrevbBlock)
	assert.Equal(t, h.Timestamp, hDecode.Timestamp)
	assert.Equal(t, h.Height, hDecode.Height)
	assert.Equal(t, h.Nonce, hDecode.Nonce)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:    1,
			PrevbBlock: types.RandomHash(),
			Timestamp:  time.Now().UnixNano(),
			Height:     10,
			Nonce:      99234,
		},
		Transactions: []Transaction{},
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b.Header.Version, bDecode.Header.Version)
	assert.Equal(t, b.Header.PrevbBlock, bDecode.Header.PrevbBlock)
	assert.Equal(t, b.Header.Timestamp, bDecode.Header.Timestamp)
	assert.Equal(t, b.Header.Height, bDecode.Header.Height)
	assert.Equal(t, b.Header.Nonce, bDecode.Header.Nonce)
}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:    1,
			PrevbBlock: types.RandomHash(),
			Timestamp:  time.Now().UnixNano(),
			Height:     10,
			Nonce:      99234,
		},
		Transactions: []Transaction{},
	}
	h := b.Hash()
	assert.False(t, h.IsZero())

}
