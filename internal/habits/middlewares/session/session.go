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
func GetSessionFromContext(req *http.Request) *sessions.Session {
	if session, ok := req.Context().Value(sessionContextKey).(*sessions.CookieStore); ok {
		store, _ := session.Get(req, "habits")
		return store
	}
	panic("Unable to obtain session from context")
}

// DeleteSessionInContext deletes the session
func DeleteSessionInContext(res http.ResponseWriter, req *http.Request) {
	if session, ok := req.Context().Value(sessionContextKey).(*sessions.CookieStore); ok {
		sess, _ := session.Get(req, "habits")
		sess.Options.MaxAge = -1
		sess.Save(req, res)

		sess, _ = session.Get(req, "current_user")
		sess.Options.MaxAge = -1
		sess.Save(req, res)
	}
}
