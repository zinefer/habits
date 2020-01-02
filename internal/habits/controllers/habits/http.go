package habits

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/zinefer/habits/internal/habits/models/habit"
)

// HabitRequest is the request payload for the Habit data model
type HabitRequest struct {
	*habit.Habit
}

// Bind a habitrequest from a http.Request
func (a *HabitRequest) Bind(r *http.Request) error {
	return nil
}

// HabitResponse is the response payload for the Habit data model
type HabitResponse struct {
	*habit.Habit
}

// NewHabitResponse returns a new HabitResponse
func NewHabitResponse(habit *habit.Habit) *HabitResponse {
	return &HabitResponse{Habit: habit}
}

// Render a HabitResponse
func (hr *HabitResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewHabitListResponse returns a new NewHabitListResponse
func NewHabitListResponse(habits []*habit.Habit) []render.Renderer {
	list := []render.Renderer{}
	for _, habit := range habits {
		list = append(list, NewHabitResponse(habit))
	}
	return list
}
