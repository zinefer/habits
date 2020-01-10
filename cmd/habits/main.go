package main

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/tasks/db"
	"github.com/zinefer/habits/internal/habits/tasks/secret"
	"github.com/zinefer/habits/internal/habits/tasks/serve"
)

var taskArg string

func main() {
	if len(os.Args) > 1 {
		taskArg = os.Args[1]

		// Remove subcommand arg
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	configuration := config.New()

	database, err := sqlx.Open("postgres", configuration.Database.URI())
	if err != nil {
		panic(err)
	}
	defer database.Close()

	subcommands := subcommander.New()
	subcommands.Register("serve", "Serve habits", serve.New(configuration, database))
	subcommands.Register("db", "Database management", db.New(configuration, database))
	subcommands.Register("secret", "Secret management", secret.New(configuration))
	success := subcommands.Execute(taskArg)

	if success {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
