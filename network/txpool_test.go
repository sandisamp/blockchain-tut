package network

import (
	"testing"

	"github.com/sandisamp/blockchain-tut/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, 0, p.Len())
}

func TestTxPoolAddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("foo"))
	assert.Nil(t, p.Add(tx), "Should not return an error when adding a valid transaction")
	assert.Equal(t, 1, p.Len())

	_ = core.NewTransaction([]byte("foo"))
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, 0, p.Len())
}
