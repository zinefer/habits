package rollback

import (
	"fmt"

	"github.com/zinefer/habits/internal/pkg/subcommander"
)

type Subcommand struct{}

// Subcommander configures the subcommander instance for this subtask
func (*Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the db subcommand
func (*Subcommand) Run() bool {
	fmt.Printf("**RUN** DB:MIGRATE:ROLLBACK")
	return true
}
