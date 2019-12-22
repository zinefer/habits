package authorize

import (
	"net/http"

	"github.com/zinefer/habits/internal/habits/middlewares/session"
	"github.com/zinefer/habits/internal/habits/models/user"
)

// AuthorizeMiddleware is a middleware that adds a database struct to the context
func AuthorizeMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session := session.GetSessionFromContext(r.Context())
			store, _ := session.Get(r, "habits")
			if user, _ := store.Values["current_user"].(*user.User); user == nil {
				http.Redirect(w, r, "/", 302)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
