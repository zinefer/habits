package generate

import (
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
)

// Subcommand for the schema task group
type Subcommand struct {
	config *config.Configuration
}

// New secret:generate subcommand
func New(config *config.Configuration) *Subcommand {
	return &Subcommand{
		config: config,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (c *Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the secret:generate subcommand
func (c *Subcommand) Run() bool {
	data := c.config.ResetSecretConfig()
	return len(data) > 0
}
