package schema

import (
	"github.com/jmoiron/sqlx"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/tasks/db/schema/dump"
	"github.com/zinefer/habits/internal/habits/tasks/db/schema/load"
)

// Subcommand for the schema task group
type Subcommand struct {
	config *config.Configuration
	db     *sqlx.DB
}

// New db subcommand
func New(config *config.Configuration, db *sqlx.DB) *Subcommand {
	return &Subcommand{
		config: config,
		db:     db,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (c *Subcommand) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Hide = true
	sc.Register("load", "Recreates the database from the schema.sql file", load.New(c.config, c.db))
	sc.Register("dump", "Dumps the current environment’s schema to database/schema.sql", dump.New(c.config, c.db))
	return sc
}

// Run the db subcommand
func (c *Subcommand) Run() bool {
	sc := c.Subcommander()
	sc.PrintAvailableCommands("")
	return true
}
