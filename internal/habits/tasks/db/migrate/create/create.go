package create

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/pkg/database/migrator"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
)

// Subcommand for the db:migrate:create task group
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
func (*Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the db:migrate:create subcommand
func (c *Subcommand) Run() bool {
	if len(os.Args) < 2 {
		fmt.Println("Error: Must pass a name for the migration")
		return false
	}
	name := os.Args[1]
	migrations := migrator.New(c.db, config.DatabaseMigrationPath)
	version, err := migrations.StubNewMigration(name)
	success := (err == nil)
	if success {
		fmt.Printf("Created migration version %v", version)
	} else {
		fmt.Printf("Error creating migration version %v", version)
		fmt.Println(err)
	}
	return success
}
