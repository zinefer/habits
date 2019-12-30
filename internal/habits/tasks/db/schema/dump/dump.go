package dump

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/zinefer/habits/internal/pkg/subcommander"

	"github.com/zinefer/habits/internal/habits/config"
)

// Subcommand for the db:schema:dump task
type Subcommand struct {
	config *config.Configuration
	db     *sqlx.DB
}

// New db:schema:dump subcommand
func New(config *config.Configuration, db *sqlx.DB) *Subcommand {
	return &Subcommand{
		config: config,
		db:     db,
	}
}

// Subcommander configures the subcommander instance for this subtask
func (c *Subcommand) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

// Run the db:schema:dump subcommand
func (c *Subcommand) Run() bool {
	bytes, err := exec.Command("pg_dump", "-s", c.config.Database.URI()).Output()
	if err != nil {
		fmt.Println(err)
		return false
	}

	dump := stripComment(string(bytes))
	dump = strings.Replace(dump, stripSQL, "", 1)
	dump = removeDoubleNewline(dump)

	out := []byte(dump)
	err = ioutil.WriteFile(config.DatabaseSchemaPath, out, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}

func stripComment(source string) string {
	re := regexp.MustCompile(`\n?-{2,}[^\n]*\n?`)
	return re.ReplaceAllString(source, "")
}

func removeDoubleNewline(source string) string {
	re := regexp.MustCompile(`\n\s*\n`)
	return re.ReplaceAllString(source, "")
}
