package tloghttp

import (
	"context"
)

type ctxKey struct{}

func CreateContext(ctx context.Context) context.Context {
	//if lp, ok := ctx.Value(ctxKey{}).(*tinylog.Logger); ok {
	//	if lp == 1 {
	//		return ctx
	//	}
	//}
	return context.WithValue(ctx, ctxKey{}, 1)
}
