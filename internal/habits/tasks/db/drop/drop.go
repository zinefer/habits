package drop

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/zinefer/habits/internal/habits/config"

	"github.com/zinefer/habits/internal/pkg/database/manager"
	"github.com/zinefer/habits/internal/pkg/subcommander"
)

// Subcommand for the db:drop task
type Subcommand struct {
	config *config.Configuration
	db     *sqlx.DB
}

// New db:drop subcommand
func New(config *config.Configuration, db *sqlx.DB) *Subcommand {
	return &Subcommand{
		config: config,
		db:     db,
	}
}

// Subcommander returns a configured subcommands instance
func (*Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the db:drop subcommand
func (c *Subcommand) Run() bool {
	// HACK: db:drop has a small problem in that main.go connected to the db
	// that we're trying to drop. Lets create a new temp connection to create the
	// database with.
	uri := strings.Replace(c.config.Database.URI(), "/"+c.config.Database.Name, "", -1)
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbManager := manager.New(db)
	err = dbManager.Drop(c.config.Database.Name)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}
