package routes

import (
	"github.com/go-chi/chi"

	"github.com/zinefer/habits/internal/habits/controllers/auth"
)

// Define routes for the habits app
func Define(r *chi.Mux) {
	r.Get("/auth/{provider}/callback", auth.Callback())
	r.Get("/auth/{provider}", auth.SignIn())
	r.Get("/logout", auth.SignOut())
}
