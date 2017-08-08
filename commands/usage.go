package commands

import (
	"github.com/gnumast/td/cli"
)

type UsageCommand struct{}

const USAGE_OUTPUT = `Usage:
	command [arguments]
	help [command]

Available commands:
a             Add a new item
c             Mark an item as completed
p             Mark an item as in progress
r             Revert a status to incomplete status
t             Add tags to an item
f             Find an item
ls            List all items
Global commands:
ga            Add a new item to the global list
gf            Find an item in the global list
gl            List all the items in the global list

Other commands:
version       Displays version information
config        Lists configuration options`

// Run displays the usage instruction
func (usage *UsageCommand) Run(app *cli.Application) error {
	app.CliOutput.Write(USAGE_OUTPUT)

	return nil
}

// Validate validates the data parsed, which makes no sense in this context
func (usage UsageCommand) Validate(configuration *cli.Application) error {
	return nil
}

// Parse parses the arguments, none in this case
func (usage UsageCommand) Parse(configuration *cli.Application, args []string) error {
	return nil
}

func (usage *UsageCommand) Help(app *cli.Application) {
}
