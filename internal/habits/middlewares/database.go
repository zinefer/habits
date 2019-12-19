package middlewares

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type key int
const databaseContextKey key = iota

// DbContextMiddleware is a middleware that adds a database struct to the context
func DbContextMiddleware(db *sqlx.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), databaseContextKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetDbFromContext returns a db
func GetDbFromContext(ctx context.Context) *sqlx.DB {
	if db, ok := ctx.Value(databaseContextKey).(*sqlx.DB); ok {
		return db
	}
	panic("Unable to obtain database struct from context")
}