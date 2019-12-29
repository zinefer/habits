package subcommander

// commander interface is implemented by external subcommands
type commander interface {
	Subcommander() *Subcommander
	Run() bool
}

// command describes a subcommand for subcommander
type command struct {
	name        string
	description string
	command     commander
}

func newCommand(name string, description string, cmd commander) command {
	return command{
		name:        name,
		description: description,
		command:     cmd,
	}
}
