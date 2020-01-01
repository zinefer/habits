package habit

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	habitModel "github.com/zinefer/habits/internal/habits/models/habit"
)

type habitKey int

const habitContextKey habitKey = iota

// HabitContextMiddleware is a middleware that adds the habit to the context
func HabitContextMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			habitID := chi.URLParam(r, "habit_id")

			if len(habitID) == 0 {
				habitID = chi.URLParam(r, "id")
			}

			id, err := strconv.ParseInt(habitID, 10, 64)

			habit, err := habitModel.FindByID(r.Context(), id)
			if err != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}

			ctx := context.WithValue(r.Context(), habitContextKey, habit)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetHabitFromContext returns the habit
func GetHabitFromContext(req *http.Request) *habitModel.Habit {
	if habit, ok := req.Context().Value(habitContextKey).(*habitModel.Habit); ok {
		return habit
	}
	panic("Unable to obtain habit from context")
}
