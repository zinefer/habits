package subcommander

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Subcommander struct has a list of commands
type Subcommander struct {
	commands []command
}

// New returns a pointer to a new subcommander instance
func New() *Subcommander {
	return &Subcommander{
		commands: make([]command, 0),
	}
}

// Register a subcommand
func (c *Subcommander) Register(name string, description string, subcommand commander) {
	cmd := newCommand(name, description, subcommand)
	c.commands = append(c.commands, cmd)
}

// Execute a registered command that matches cmdArg
func (c *Subcommander) Execute(cmdArg string) bool {
	command, subcommand := splitSubtask(cmdArg)

	for _, cmd := range c.commands {
		if cmd.name == command {
			if len(subcommand) == 0 {
				return cmd.command.Run()
			}

			return cmd.command.Subcommander().Execute(subcommand)
		}
	}

	c.PrintAvailableCommands(command)
	return true
}

// PrintAvailableCommands prints registered subcommands and descriptions
func (c *Subcommander) PrintAvailableCommands(parent string) {
	if len(c.commands) > 0 {
		fmt.Printf("Available subcommands:")
		w := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)
		printCommandsOnCommander(w, parent, c)
		w.Flush()
	}
}

func printCommandsOnCommander(w *tabwriter.Writer, parent string, c *Subcommander) {
	for _, sc := range c.commands {
		name := sc.name
		if len(parent) > 0 {
			name = parent + ":" + name
		}
		fmt.Fprintf(w, "\n\t%v\t%v", name, sc.description)
		printCommandsOnCommander(w, name, sc.command.Subcommander())
	}
}

// splitSubtask splits taskArgs into (command, subcommand)
func splitSubtask(taskArg string) (string, string) {
	tokens := strings.SplitN(taskArg, ":", 2)
	if len(tokens) == 1 {
		return tokens[0], ""
	}
	return tokens[0], tokens[1]
}
