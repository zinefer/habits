package routes

import (
	"github.com/go-chi/chi"

	"github.com/zinefer/habits/internal/habits/middlewares/authorize"
	"github.com/zinefer/habits/internal/habits/middlewares/habit"

	"github.com/zinefer/habits/internal/habits/controllers/activities"
	"github.com/zinefer/habits/internal/habits/controllers/auth"
	"github.com/zinefer/habits/internal/habits/controllers/habits"
)

// Define routes for the habits app
func Define(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {
		r.Get("/auth/{provider}/callback", auth.Callback())
		r.Get("/auth/{provider}", auth.SignIn())
		r.Get("/logout", auth.SignOut())

		r.Get("/habits/{user:[a-z0-9-_]+}", habits.UserList())

		r.Route("/habits", func(r chi.Router) {
			r.Use(authorize.AuthorizeMiddleware())

			r.Get("/", habits.List())
			r.Post("/", habits.Create())
			r.Route("/{id:[0-9]+}", func(r chi.Router) {
				r.Use(habit.HabitContextMiddleware())

				r.Get("/", habits.Show())

				rHO := r.With(habit.HabitOwnerMiddleware())
				rHO.Patch("/", habits.Update())
				rHO.Delete("/", habits.Delete())
			})
			r.Route("/{habit_id:[0-9]+}/activities", func(r chi.Router) {
				r.Use(habit.HabitContextMiddleware())

				r.Get("/", activities.ListLastYear())
				r.Get("/streaks", activities.Streaks())

				rHO := r.With(habit.HabitOwnerMiddleware())
				rHO.Post("/", activities.Create())
				rHO.Delete("/{id:[0-9]+}", activities.Delete())
			})
		})
	})
}
