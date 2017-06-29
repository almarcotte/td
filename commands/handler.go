package commands

import (
	"fmt"
)

type Configuration struct {
}

// Runner represents a type that can be executed. All commands should implement this.
type Runner interface {
	Run(*Configuration) error
}

// Validator represents a type that can be validated. All commands should implement this.
type Validator interface {
	Validate(*Configuration) error
}

// Handle accepts everything passed via the command line and figures out which command to call
func Handle(args []string) (err error) {
	cmd, opts := args[0], args[1:]

	fmt.Printf("Not implemented yet. Received: %v, Options: %v\n", cmd, opts)

	return
}
