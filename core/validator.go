package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("Block with height %d already exists with hash %s", b.Height, b.Hash(BlockHasher{}))
	}

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("Block height is %d, expected %d", b.Height, v.bc.Height()+1)
	}
	prevHeader, err := v.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)
	if b.PrevbBlockHash != hash {
		return fmt.Errorf("Block prev hash is %s which is invalid", b.PrevbBlockHash)
	}

	if err := b.Verify(); err != nil {
		return err
	}
	return nil
}
