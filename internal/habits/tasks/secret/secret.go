package secret

import (
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
	"github.com/zinefer/habits/internal/habits/tasks/secret/generate"
)

// Subcommand for the schema task group
type Subcommand struct {
	config *config.Configuration
}

// New secret subcommand group
func New(config *config.Configuration) *Subcommand {
	return &Subcommand{
		config: config,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (c *Subcommand) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Hide = true
	sc.Register("generate", "Generates a secret session key", generate.New(c.config))
	return sc
}

// Run the secret subcommand
func (c *Subcommand) Run() bool {
	sc := c.Subcommander()
	sc.PrintAvailableCommands("")
	return true
}
