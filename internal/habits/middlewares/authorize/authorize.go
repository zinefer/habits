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
			sess := session.GetSessionFromContext(r)
			if user, _ := sess.Values["current_user"].(*user.User); user == nil {
				session.DeleteSessionInContext(w, r)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
