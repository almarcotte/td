package commands

import "github.com/gnumast/td/todo"

// Command represents a supported action that can parse command line arguments, validate them and run.
type Command interface {
	Run(*todo.Configuration) error
	Validate(*todo.Configuration) error
	Parse(*todo.Configuration, []string) error
}
