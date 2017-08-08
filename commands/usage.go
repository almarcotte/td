package commands

import (
	"fmt"
	"github.com/gnumast/td/cli"
)

type UsageCommand struct{}

// Run displays the usage instruction
func (usage *UsageCommand) Run(app *cli.Application) error {
	fmt.Println("Usage:\n  command [arguments]\n  command help")

	fmt.Println()

	fmt.Println("Available commands:")
	fmt.Println("  a             Add a new item")
	fmt.Println("  c             Mark an item as completed")
	fmt.Println("  p             Mark an item as in progress")
	fmt.Println("  r             Revert a status to incomplete status")
	fmt.Println("  t             Add tags to an item")
	fmt.Println("  f             Find an item")
	fmt.Println("  ls            List all items")

	fmt.Println()

	fmt.Println("Global commands:")
	fmt.Println("  ga            Add a new item to the global list")
	fmt.Println("  gf            Find an item in the global list")
	fmt.Println("  gl            List all the items in the global list")

	fmt.Println()

	fmt.Println("Other commands:")
	fmt.Println("  version       Displays version information")
	fmt.Println("  config        Lists configuration options")

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

func (usage UsageCommand) Help() string {
	return ""
}
