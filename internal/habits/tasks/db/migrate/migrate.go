package migrate

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/pkg/database/migrator"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/tasks/db/migrate/create"
	"github.com/zinefer/habits/internal/habits/tasks/db/migrate/rollback"
	"github.com/zinefer/habits/internal/habits/tasks/db/schema/dump"
)

// Subcommand for the db:migrate task group
type Subcommand struct {
	config *config.Configuration
	db     *sqlx.DB
}

// New db:migrate subcommand
func New(config *config.Configuration, db *sqlx.DB) *Subcommand {
	return &Subcommand{
		config: config,
		db:     db,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (c *Subcommand) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Register("rollback", "Rolls back the last migration", &rollback.Subcommand{})
	sc.Register("create", "Stubs out a new database migration", create.New(c.config, c.db))
	return sc
}

// Run the db:migrate subcommand
func (c *Subcommand) Run() bool {
	migrations := migrator.New(c.db, config.DatabaseMigrationPath)
	err := migrations.Migrate()
	if err != nil {
		fmt.Println(err)
	}

	d := dump.New(c.config, c.db)
	dumped := d.Run()

	return err == nil && dumped
}
