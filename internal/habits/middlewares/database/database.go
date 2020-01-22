package database

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type dbKey int

const databaseContextKey dbKey = iota

// DbContextMiddleware is a middleware that adds a database struct to the context
func DbContextMiddleware(db *sqlx.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := SetDbInContext(r.Context(), db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// SetDbInContext returns a copy of ctx with a db set
func SetDbInContext(ctx context.Context, db *sqlx.DB) context.Context {
	return context.WithValue(ctx, databaseContextKey, db)
}

// GetDbFromContext returns a db
func GetDbFromContext(ctx context.Context) *sqlx.DB {
	if db, ok := ctx.Value(databaseContextKey).(*sqlx.DB); ok {
		return db
	}
	panic("Unable to obtain database struct from context")
}
