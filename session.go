package gearsession

import (
	"context"

	"github.com/go-session/session"
	"github.com/teambition/gear"
)

// specify the context key
type (
	sessionKey struct{}
	manageKey  struct{}
)

// New create a session middleware
func New(opt ...session.Option) gear.Middleware {
	manage := session.NewManager(opt...)
	return func(ctx *gear.Context) error {
		ctx.SetAny(manageKey{}, manage)
		store, err := manage.Start(context.Background(), ctx.Res, ctx.Req)
		if err != nil {
			return err
		}
		ctx.SetAny(sessionKey{}, store)
		return nil
	}
}

// FromContext get session storage from context
func FromContext(ctx *gear.Context) session.Store {
	return ctx.MustAny(sessionKey{}).(session.Store)
}

// Destroy a session
func Destroy(ctx *gear.Context) error {
	return ctx.MustAny(manageKey{}).(*session.Manager).Destroy(context.Background(), ctx.Res, ctx.Req)
}

// Refresh a session and return to session storage
func Refresh(ctx *gear.Context) (session.Store, error) {
	return ctx.MustAny(manageKey{}).(*session.Manager).Refresh(context.Background(), ctx.Res, ctx.Req)
}
