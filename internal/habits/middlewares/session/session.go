package session

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

type sessKey int

const sessionContextKey sessKey = iota

// SessionContextMiddleware is a middleware that adds a database struct to the context
func SessionContextMiddleware(session *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), sessionContextKey, session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetSessionFromContext returns the session
func GetSessionFromContext(ctx context.Context) *sessions.CookieStore {
	if session, ok := ctx.Value(sessionContextKey).(*sessions.CookieStore); ok {
		return session
	}
	panic("Unable to obtain session from context")
}
