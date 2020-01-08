package habit

import (
	"context"
	"time"

	"github.com/zinefer/habits/internal/habits/middlewares/database"

	"github.com/zinefer/habits/internal/habits/models/activity"
)

// Habit model
type Habit struct {
	ID      int64
	UserID  int64 `db:"user_id"`
	Name    string
	Created time.Time
}

// New Creates a Habit model
func New(userID int64, name string) *Habit {
	return &Habit{
		UserID: userID,
		Name:   name,
	}
}

// Save a Habit to the database
func (h *Habit) Save(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	stmt, err := db.PrepareNamed("INSERT INTO habits (user_id, name, created) VALUES (:user_id, :name, :created) RETURNING id;")
	if err != nil {
		return err
	}
	return stmt.Get(&h.ID, h)
}

// Update a Habit in the database
func (h *Habit) Update(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	stmt, err := db.PrepareNamed("UPDATE habits SET user_id = :user_id, name = :name, created = :created WHERE id = :id;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(h)
	return err
}

// Delete a Habit from the database
func (h *Habit) Delete(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	err := activity.DeleteAllByHabit(ctx, h.ID)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM habits WHERE id = $1;", h.ID)
	return err
}

// CountActivitiesInLastYear counts the activities in the past year
func (h *Habit) CountActivitiesInLastYear(ctx context.Context) ([]*activity.ActivityCount, error) {
	return activity.CountByDayInLastYearByHabit(ctx, h.ID)
}

// GetStreaks returns activity streaks
func (h *Habit) GetStreaks(ctx context.Context) (*activity.ActivityStreaks, error) {
	return activity.GetStreaksByHabit(ctx, h.ID)
}

// GetActivities returns a list of habits for a user
func (h *Habit) GetActivities(ctx context.Context) ([]*activity.Activity, error) {
	return activity.FindAllByHabit(ctx, h.ID)
}

// FindAllByUser returns a list of habits owned by a user
func FindAllByUser(ctx context.Context, userID int64) ([]*Habit, error) {
	db := database.GetDbFromContext(ctx)
	habits := []*Habit{}
	err := db.Select(&habits, "SELECT * FROM habits WHERE user_id = $1;", userID)
	return habits, err
}

// FindByID returns a habit by it's ID
func FindByID(ctx context.Context, habitID int64) (*Habit, error) {
	habit := &Habit{}
	db := database.GetDbFromContext(ctx)
	err := db.Get(habit, "SELECT * FROM habits WHERE id = $1 LIMIT 1;", habitID)
	return habit, err
}
