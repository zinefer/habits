package activities

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/zinefer/habits/internal/habits/models/activity"
)

// ActivityRequest is the request payload for the Activity data model
type ActivityRequest struct {
	*activity.Activity
}

// Bind a ActivityRequest from a http.Request
func (a *ActivityRequest) Bind(r *http.Request) error {
	return nil
}

// ActivityCountResponse is the response payload for the Activity data model
type ActivityCountResponse struct {
	*activity.ActivityCount
}

// NewActivityCountResponse returns a new ActivityResponse
func NewActivityCountResponse(activity *activity.ActivityCount) *ActivityCountResponse {
	return &ActivityCountResponse{ActivityCount: activity}
}

// Render a ActivityCountResponse
func (hr *ActivityCountResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewActivityCountListResponse returns a new NewActivityListResponse
func NewActivityCountListResponse(activitys []*activity.ActivityCount) []render.Renderer {
	list := []render.Renderer{}
	for _, activity := range activitys {
		list = append(list, NewActivityCountResponse(activity))
	}
	return list
}
