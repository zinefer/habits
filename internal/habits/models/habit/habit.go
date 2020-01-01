package habit

import (
	"context"
	"time"

	"github.com/zinefer/habits/internal/habits/middlewares/database"
)

// Habit model
type Habit struct {
	ID      int64
	UserID  int64 `db:"user_id"`
	Created time.Time
}

// New Creates a Habit model
func New(userID int64) *Habit {
	return &Habit{
		UserID: userID,
	}
}

// Save a Habit to the database
func (h *Habit) Save(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	stmt, err := db.PrepareNamed("INSERT INTO habits (user_id, created) VALUES (:user_id, :created) RETURNING id")
	if err != nil {
		return err
	}
	return stmt.Get(&h.ID, h)
}
