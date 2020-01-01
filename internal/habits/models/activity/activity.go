package activity

import (
	"context"
	"time"

	"github.com/zinefer/habits/internal/habits/middlewares/database"
)

// Activity model
type Activity struct {
	ID      int64
	HabitID int64 `db:"habit_id"`
	Created time.Time
}

// New Activity model
func New(habitID int64) *Activity {
	return &Activity{
		HabitID: habitID,
	}
}

// Save an Activity to the database
func (a *Activity) Save(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	stmt, err := db.PrepareNamed("INSERT INTO activities (habit_id) VALUES (:habit_id) RETURNING id")
	if err != nil {
		return err
	}
	return stmt.Get(&a.ID, a)
}
