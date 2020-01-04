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
	stmt, err := db.PrepareNamed("INSERT INTO activities (habit_id) VALUES (:habit_id) RETURNING id;")
	if err != nil {
		return err
	}
	return stmt.Get(&a.ID, a)
}

// FindAllByHabit returns a list of activities owned by a habit
func FindAllByHabit(ctx context.Context, habitID int64) ([]*Activity, error) {
	db := database.GetDbFromContext(ctx)
	activities := []*Activity{}
	err := db.Select(&activities, "SELECT * FROM activities WHERE habit_id = $1;", habitID)
	return activities, err
}

// ActivityCount model
type ActivityCount struct {
	Day   time.Time
	Count int
}

// CountByDayInLastYearByHabit returns a count of the last years activity by habit
func CountByDayInLastYearByHabit(ctx context.Context, habitID int64) ([]*ActivityCount, error) {
	db := database.GetDbFromContext(ctx)
	counts := []*ActivityCount{}
	err := db.Select(&counts, `WITH date_range AS (
		SELECT date_trunc('day', now() - (364 + extract(dow from now())) * interval '1 day') as beginning,
			   date_trunc('day', now()) as ending
	)
	SELECT day, sum(count) as count FROM (
		SELECT date_trunc('day', activities.created)::date as "day", count(*) as count
		FROM activities
		WHERE activities.created >= (select beginning from date_range) AND habit_id = $1
		GROUP BY day
		UNION
		SELECT day::date, 0 as count
		FROM generate_series((select beginning from date_range), (select ending from date_range), '1 day') day
	) as activities GROUP BY day ORDER BY day`, habitID)
	return counts, err
}
