package commands

import (
	"fmt"
)



type UsageCommand struct{}

// Run displays the usage instruction
func (usage *UsageCommand) Run(configuration *Configuration) error {
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

func (usage *UsageCommand) Validate(configuration *Configuration) error {
	return nil
}

func (usage *UsageCommand) Parse(configuration *Configuration, args []string) error {
	return nil
}
