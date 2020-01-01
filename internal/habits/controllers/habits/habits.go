package habits

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/zinefer/habits/internal/habits/middlewares/session"

	"github.com/zinefer/habits/internal/habits/models/habit"
	"github.com/zinefer/habits/internal/habits/models/user"
)

// Create a habit
func Create() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		data := &HabitRequest{}
		if err := render.Bind(req, data); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}

		sess := session.GetSessionFromContext(req)
		currentUser := sess.Values["current_user"].(*user.User)

		habit := habit.New(currentUser.ID)
		habit.Save(req.Context())

		render.Status(req, http.StatusCreated)
		render.Render(res, req, NewHabitResponse(habit))
	}
}

// Show a habit
func Show() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

// List habits
func List() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

// Update a habit
func Update() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

// Delete a habit
func Delete() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
