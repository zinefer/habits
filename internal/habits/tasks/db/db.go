package db

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jmoiron/sqlx"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/tasks/db/create"
	"github.com/zinefer/habits/internal/habits/tasks/db/drop"
	"github.com/zinefer/habits/internal/habits/tasks/db/migrate"
	"github.com/zinefer/habits/internal/habits/tasks/db/schema"
)

// Subcommand for the db task
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
	sc.Register("create", "Creates the database for current environment", create.New(c.config, c.db))
	sc.Register("drop", "Drops the database for current environment", drop.New(c.config, c.db))
	sc.Register("migrate", "Runs migrations for the current environment that have not run yet", migrate.New(c.db))
	sc.Register("schema", "Database schema subcommand group", schema.New(c.config, c.db))
	return sc
}

// Run the db subcommand
func (c *Subcommand) Run() bool {
	fmt.Printf("Use \\q to exit\n")
	cmd := exec.Command("psql", c.config.Database.URI())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	return err == nil
}
