package main

import (
	"os"

	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/tasks/db"
	"github.com/zinefer/habits/internal/habits/tasks/serve"
)

var taskArg string

func main() {
	if len(os.Args) > 1 {
		taskArg = os.Args[1]

		// Remove subcommand arg
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	subcommands := subcommander.New()
	subcommands.Register("serve", "Serve habits", &serve.Subcommand{})
	subcommands.Register("db", "Database management", &db.Subcommand{})
	subcommands.Execute(taskArg)
}
