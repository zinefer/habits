package activities

import (
	"net/http"

	"github.com/go-chi/render"

	habitMW "github.com/zinefer/habits/internal/habits/middlewares/habit"
)

// Create an activity for a habit
func Create() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

// ListLastYear of activities for a habit
func ListLastYear() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		habit := habitMW.GetHabitFromContext(req)
		activities, err := habit.CountActivitiesInLastYear(req.Context())
		if err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}

		if err := render.RenderList(res, req, NewActivityCountListResponse(activities)); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}
	}
}

// Delete an activity for a habit
func Delete() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
