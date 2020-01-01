package routes

import (
	"github.com/go-chi/chi"

	"github.com/zinefer/habits/internal/habits/middlewares/authorize"
	"github.com/zinefer/habits/internal/habits/middlewares/habit"

	"github.com/zinefer/habits/internal/habits/controllers/auth"
	"github.com/zinefer/habits/internal/habits/controllers/habits"
	"github.com/zinefer/habits/internal/habits/controllers/activities"
)

// Define routes for the habits app
func Define(r *chi.Mux) {
	r.Get("/auth/{provider}/callback", auth.Callback())
	r.Get("/auth/{provider}", auth.SignIn())
	r.Get("/logout", auth.SignOut())

	r.Route("/habits", func(r chi.Router) {
		r.Use(authorize.AuthorizeMiddleware())

		r.Get("/{id}", habits.Show())
		r.Get("/", habits.List())
		r.Post("/", habits.Create())
		r.Patch("/{id}", habits.Update())
		r.Delete("/{id}", habits.Delete())

		r.Route("/{habit_id}/activities", func(r chi.Router) {
			r.Use(habit.HabitContextMiddleware())

			r.Get("/", activities.List())
			r.Post("/", activities.Create())
			r.Delete("/{id}", activities.Delete())
		})
	})
}
