package valuecontext

import (
	"context"
	"log/slog"
	"sync"
)

type ValueKeyType string

const (
	CtxKey      ValueKeyType = "CTX-KEY"
	CtxValueKey ValueKeyType = "CTX-VALUE-KEY"
	IsTxError   ValueKeyType = "IS-TX-ERROR"
)

type Committer interface {
	Begin() Committer
	Commit() error
	Rollback() error
	Tx() any
}

type ContextValue struct {
	Tx     Committer
	Logger *slog.Logger
}

func NewValueContext(parent context.Context, val *ContextValue) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	m := new(sync.Map)

	m.Store(CtxValueKey, val)

	return context.WithValue(parent, CtxKey, m)
}

func tryGetValueFromContext(ctx context.Context) (*ContextValue, bool) {
	val := ctx.Value(CtxKey)

	if val == nil {
		return nil, false
	}

	m, ok := val.(*sync.Map)
	if !ok {
		return nil, false
	}

	v, ok := m.Load(CtxValueKey)
	if !ok {
		return nil, false
	}

	ctxVal, ok := v.(*ContextValue)
	if !ok || ctxVal == nil {
		return nil, false
	}

	return ctxVal, true
}

func TryGetTxFromContext(ctx context.Context) (Committer, bool) {
	ctxVal, ok := tryGetValueFromContext(ctx)
	if !ok {
		return nil, false
	}

	return ctxVal.Tx, true
}

func GetLogger(ctx context.Context) *slog.Logger {
	val, _ := tryGetValueFromContext(ctx)
	return val.Logger
}

func SetTx(ctx context.Context, tx Committer) {
	val, ok := tryGetValueFromContext(ctx)
	if !ok {
		return
	}

	val.Tx = tx
}
