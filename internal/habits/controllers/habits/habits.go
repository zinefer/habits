package habits

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	habitMW "github.com/zinefer/habits/internal/habits/middlewares/habit"
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

		habit := habit.New(currentUser.ID, data.Name)
		if err := habit.Save(req.Context()); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}

		render.Status(req, http.StatusCreated)
		render.Render(res, req, NewHabitResponse(habit))
	}
}

// Show a habit
func Show() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		data := &HabitRequest{}
		if err := render.Bind(req, data); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}

		habit := habitMW.GetHabitFromContext(req)

		render.Status(req, http.StatusCreated)
		render.Render(res, req, NewHabitResponse(habit))
	}
}

// List habits
func List() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		sess := session.GetSessionFromContext(req)
		currentUser := sess.Values["current_user"].(*user.User)

		habits, err := currentUser.GetHabits(req.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(res, http.StatusText(400), 400)
			return
		}

		if err := render.RenderList(res, req, NewHabitListResponse(habits)); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}
	}
}

// Update a habit
func Update() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		data := &HabitRequest{}
		if err := render.Bind(req, data); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}

		habit := habitMW.GetHabitFromContext(req)

		render.Status(req, http.StatusCreated)
		render.Render(res, req, NewHabitResponse(habit))
	}
}

// Delete a habit
func Delete() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		habit := habitMW.GetHabitFromContext(req)

		err := habit.Delete(req.Context())
		if err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}

		http.Error(res, http.StatusText(200), 200)
	}
}
