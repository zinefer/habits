package create

import (
	"fmt"

	"github.com/zinefer/habits/internal/pkg/subcommander"
)

type Subcommand struct{}

// Subcommander returns a configured subcommands instance
func (*Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the db subcommand
func (*Subcommand) Run() bool {
	fmt.Printf("**RUN** DB:CREATE")
	return true
}
