package generate

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

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
	data := generateRandomKey()
	file, err := os.Create(config.SecretConfigPath)
	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}

func generateRandomKey() []byte {
	k := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}
