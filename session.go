package gearsession

import (
	"context"

	"github.com/teambition/gear"
	"gopkg.in/session.v2"
)

// Specify global session manager
var globalManager *session.Manager

// Specify the context key
type sessionKey struct{}

// New Create a session middleware
func New(opt ...session.Option) gear.Middleware {
	globalManager = session.NewManager(opt...)
	return func(ctx *gear.Context) error {
		store, err := globalManager.Start(context.Background(), ctx.Res, ctx.Req)
		if err != nil {
			return err
		}
		ctx.SetAny(sessionKey{}, store)
		return nil
	}
}

// FromContext Get session storage from context
func FromContext(ctx *gear.Context) session.Store {
	val := ctx.MustAny(sessionKey{})
	return val.(session.Store)
}

// Destroy Destroy a session
func Destroy(ctx *gear.Context) error {
	if globalManager == nil {
		return nil
	}
	return globalManager.Destroy(context.Background(), ctx.Res, ctx.Req)
}

// Refresh a session and return to session storage
func Refresh(ctx *gear.Context) (session.Store, error) {
	if globalManager == nil {
		return nil, nil
	}
	return globalManager.Refresh(context.Background(), ctx.Res, ctx.Req)
}
