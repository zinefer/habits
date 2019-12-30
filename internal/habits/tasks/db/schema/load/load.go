package load

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/pkg/database/manager"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
)

// Subcommand for the db:schema:load task
type Subcommand struct {
	config *config.Configuration
	db     *sqlx.DB
}

// New db:schema:load subcommand
func New(config *config.Configuration, db *sqlx.DB) *Subcommand {
	return &Subcommand{
		config: config,
		db:     db,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (c *Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the db:schema:load subcommand
func (c *Subcommand) Run() bool {
	dbManager := manager.New(c.db)
	err := dbManager.Load(config.DatabaseSchemaPath)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}
