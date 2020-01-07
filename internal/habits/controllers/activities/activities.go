package activities

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	habitMW "github.com/zinefer/habits/internal/habits/middlewares/habit"

	"github.com/zinefer/habits/internal/habits/models/activity"
)

// Create an activity for a habit
func Create() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		data := &ActivityRequest{}
		if err := render.Bind(req, data); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}

		habit := habitMW.GetHabitFromContext(req)

		activity := activity.New(habit.ID)
		err := activity.Save(req.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(res, http.StatusText(400), 400)
			return
		}

		render.Status(req, http.StatusCreated)
	}
}

// ListLastYear of activities for a habit
func ListLastYear() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		habit := habitMW.GetHabitFromContext(req)
		activities, err := habit.CountActivitiesInLastYear(req.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(res, http.StatusText(400), 400)
			return
		}

		if err := render.RenderList(res, req, NewActivityCountListResponse(activities)); err != nil {
			http.Error(res, http.StatusText(400), 400)
			return
		}
	}
}

// Streaks of activities for a habit
func Streaks() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		habit := habitMW.GetHabitFromContext(req)

		streak, err := habit.GetStreaks(req.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(res, http.StatusText(400), 400)
			return
		}

		render.Render(res, req, NewActivityStreaksResponse(streak))
	}
}

// Delete an activity for a habit
func Delete() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
