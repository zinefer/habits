package migrate

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/tasks/db/migrate/rollback"
)

// Subcommand the functionality
type Subcommand struct {
	db *sqlx.DB
}

// New db:migrate subcommand
func New(db *sqlx.DB) *Subcommand {
	return &Subcommand{
		db: db,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (*Subcommand) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Register("rollback", "Rolls back the last migration", &rollback.Subcommand{})
	return sc
}

// Run the db subcommand
func (*Subcommand) Run() bool {
	fmt.Printf("**RUN** DB:MIGRATE")
	return true
}
