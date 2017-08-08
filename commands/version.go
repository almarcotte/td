package commands

import (
	"fmt"
	"github.com/gnumast/td/todo"
	"runtime"
)

const (
	VERSION = "0.0.1"
)

type VersionCommand struct{}

func (vers VersionCommand) Run(configuration *todo.Application) error {
	fmt.Printf("td version %s %s/%s\n", VERSION, runtime.GOOS, runtime.GOARCH)

	return nil
}

func (vers VersionCommand) Parse(configuration *todo.Application, args []string) error {
	return nil
}

func (vers VersionCommand) Validate(configuration *todo.Application) error {
	return nil
}

func (vers VersionCommand) Help() string {
	return ""
}
