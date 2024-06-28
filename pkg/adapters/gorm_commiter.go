package adapter

import (
	"service/pkg/valuecontext"

	"gorm.io/gorm"
)

type GormCommiter struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewGormCommiter(db *gorm.DB) valuecontext.Committer {
	return &GormCommiter{db: db}
}

func (c *GormCommiter) Tx() any {
	return c.tx
}

func (c *GormCommiter) Begin() valuecontext.Committer {
	tx := c.db.Begin()
	c.tx = tx
	return c
}

func (c *GormCommiter) Commit() error {
	if c.tx == nil {
		return nil
	}
	return c.tx.Commit().Error
}

func (c *GormCommiter) Rollback() error {
	if c.tx == nil {
		return nil
	}
	return c.tx.Rollback().Error
}
