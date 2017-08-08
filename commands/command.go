package commands

import "github.com/gnumast/td/todo"

// Command represents a supported action that can parse command line arguments, validate them and run.
type Command interface {
	Run(*todo.Application) error
	Validate(*todo.Application) error
	Parse(*todo.Application, []string) error
	Help() string
}
