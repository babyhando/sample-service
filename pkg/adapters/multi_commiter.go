package adapter

import (
	"service/pkg/valuecontext"
)

type multiCommiter struct {
	commiters []valuecontext.Committer
	txs       []valuecontext.Committer
}

func NewMultiCommiter(commiters []valuecontext.Committer) *multiCommiter {
	mc := &multiCommiter{
		commiters: commiters,
	}

	txs := make([]valuecontext.Committer, len(commiters))
	mc.txs = txs

	return mc
}

func (c *multiCommiter) Begin() valuecontext.Committer {
	for i, cm := range c.commiters {
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
