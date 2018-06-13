package gearsession

import (
	"context"
	"sync"

	"github.com/go-session/session"
	"github.com/teambition/gear"
)

var (
	once            sync.Once
	internalManager *session.Manager
)

// specify the context key
type sessionKey struct{}

func manager(opt ...session.Option) *session.Manager {
	once.Do(func() {
		internalManager = session.NewManager(opt...)
	})
	return internalManager
}

// New create a session middleware
func New(opt ...session.Option) gear.Middleware {
	return func(ctx *gear.Context) error {
		store, err := manager(opt...).Start(context.Background(), ctx.Res, ctx.Req)
		if err != nil {
			return err
		}
		ctx.SetAny(sessionKey{}, store)
		return nil
	}
}

// FromContext get session storage from context
func FromContext(ctx *gear.Context) session.Store {
	val := ctx.MustAny(sessionKey{})
	return val.(session.Store)
}

// Destroy a session
func Destroy(ctx *gear.Context) error {
	return manager().Destroy(context.Background(), ctx.Res, ctx.Req)
}

// Refresh a session and return to session storage
func Refresh(ctx *gear.Context) (session.Store, error) {
	return manager().Refresh(context.Background(), ctx.Res, ctx.Req)
}
