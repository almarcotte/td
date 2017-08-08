package commands

import (
	"github.com/gnumast/td/cli"
)

// Command represents a supported action that can parse command line arguments, validate them and run.
type Command interface {
	Run(*cli.Application) error
	Validate(*cli.Application) error
	Parse(*cli.Application, []string) error
	Help(*cli.Application)
}
