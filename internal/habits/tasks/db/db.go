package db

import (
	"fmt"

	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/tasks/db/create"
	"github.com/zinefer/habits/internal/habits/tasks/db/migrate"
)

// Subcommand for the db task
type Subcommand struct{}

// Subcommander configures the subcommander instance for this subtask
func (*Subcommand) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Register("create", "Create database for application", &create.Subcommand{})
	sc.Register("migrate", "Runs migrations for the current environment that have not run yet", &migrate.Subcommand{})
	return sc
}

// Run the db subcommand
func (c *Subcommand) Run() bool {
	fmt.Printf("**RUN** DB")
	return true
}
