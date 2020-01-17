package activity

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/habits/middlewares/database"
)

// Activity model
type Activity struct {
	ID       int64
	HabitID  int64 `db:"habit_id"`
	Moment   time.Time
	TimeZone int64 `db:"time_zone"`
	Created  time.Time
}

// New Activity model
func New(habitID int64, moment time.Time, timezone int64) *Activity {
	return &Activity{
		HabitID:  habitID,
		Moment: moment,
		TimeZone: timezone,
	}
}

// Save an Activity to the database
func (a *Activity) Save(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	var stmt *sqlx.NamedStmt
	var err error
	if a.Moment.IsZero() {
		stmt, err = db.PrepareNamed("INSERT INTO activities (habit_id, moment, time_zone) VALUES (:habit_id, CURRENT_TIMESTAMP + :time_zone * INTERVAL '1 hour', :time_zone) RETURNING id;")
	} else {
		stmt, err = db.PrepareNamed("INSERT INTO activities (habit_id, moment, time_zone) VALUES (:habit_id, :moment, :time_zone) RETURNING id;")		
	}
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

// DeleteAllByHabit removes activities from a habit
func DeleteAllByHabit(ctx context.Context, habitID int64) error {
	db := database.GetDbFromContext(ctx)
	_, err := db.Exec("DELETE FROM activities WHERE habit_id = $1;", habitID)
	return err
}

// ActivityCount model
type ActivityCount struct {
	Day   time.Time
	Count int
}

// CountByDayInLastYearByHabit returns a count of the last years activity by habit
func CountByDayInLastYearByHabit(ctx context.Context, habitID int64, timezone int) ([]*ActivityCount, error) {
	db := database.GetDbFromContext(ctx)
	counts := []*ActivityCount{}
	err := db.Select(&counts, `WITH now AS (
		SELECT CURRENT_TIMESTAMP + $2 * INTERVAL '1 hour' AS timestamp
	), date_range AS (
	  SELECT date_trunc('day', (SELECT timestamp FROM now) - 371 * interval '1 day') as beginning,
	  date_trunc('day', (SELECT timestamp FROM now)) as ending
	)	
	SELECT day, sum(count) as count FROM (
		SELECT date_trunc('day', activities.moment)::date as "day", count(*) as count
		FROM activities
		WHERE activities.moment >= (select beginning from date_range) AND habit_id = $1
		GROUP BY day
	) as activities GROUP BY day ORDER BY day;`, habitID, timezone)
	return counts, err
}

// ActivityStreak model
type ActivityStreak struct {
	Streak  int
	MinDate time.Time
	MaxDate time.Time
}

// ActivityStreaks model
type ActivityStreaks struct {
	Longest *ActivityStreak
	Current *ActivityStreak
}

// GetStreaksByHabit returns a habits streaks
func GetStreaksByHabit(ctx context.Context, habitID int64) (*ActivityStreaks, error) {
	db := database.GetDbFromContext(ctx)
	streaks := &ActivityStreaks{}
	tx, err := db.Begin()
	if err != nil {
		return streaks, err
	}

	_, err = tx.Exec(`WITH groups(date, grp) AS (
		SELECT 
		  DISTINCT moment::date, 
		  EXTRACT(epoch from moment::date)::int / 86400 - DENSE_RANK() OVER (ORDER BY moment::date) AS grp
		FROM activities
		WHERE habit_id = $1
	  )
	  SELECT
		COUNT(*) AS streak,
		MIN(date) as minDate,
		MAX(date) as maxDate
	  INTO TEMPORARY TABLE streaks
	  FROM groups
	  GROUP BY grp;`, habitID)

	if err != nil {
		return streaks, err
	}

	longest, _ := makeStreakSubquery(tx, "SELECT * FROM streaks ORDER BY streak DESC, minDate DESC LIMIT 1")
	current, _ := makeStreakSubquery(tx, "SELECT * FROM streaks WHERE maxDate = now()::date LIMIT 1")
	tx.Rollback()

	streaks.Longest = longest
	streaks.Current = current
	return streaks, err
}

func makeStreakSubquery(tx *sql.Tx, query string) (*ActivityStreak, error) {
	streak := &ActivityStreak{}
	row := tx.QueryRow(query)
	err := row.Scan(&streak.Streak, &streak.MinDate, &streak.MaxDate)
	if err != nil {
		return streak, err
	}

	return streak, err
}
