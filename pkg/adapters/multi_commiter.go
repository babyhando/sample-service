package adapter

import (
	"service/pkg/valuecontext"
)

type multiCommiter struct {
	committers []valuecontext.Committer
	txs        []valuecontext.Committer
}

func NewMultiCommiter(committers []valuecontext.Committer) *multiCommiter {
	mc := &multiCommiter{
		committers: committers,
	}

	txs := make([]valuecontext.Committer, len(committers))
	mc.txs = txs

	return mc
}

func (c *multiCommiter) Begin() valuecontext.Committer {
	for i, cm := range c.committers {
		c.txs[i] = cm.Begin()
	}

	return c
}

func (c *multiCommiter) Commit() error {
	for _, tx := range c.txs {
		tx.Commit()
	}
	return nil
}

func (c *multiCommiter) Rollback() error {
	for _, tx := range c.txs {
		tx.Rollback()
	}
	return nil
}

func (c *multiCommiter) Tx() any {
	return c.txs // order matters
}
