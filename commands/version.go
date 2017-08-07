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

func (vers *VersionCommand) Run(configuration *todo.Configuration) error {
	fmt.Printf("td version %s %s/%s\n", VERSION, runtime.GOOS, runtime.GOARCH)

	return nil
}

func (vers *VersionCommand) Parse(configuration *todo.Configuration, args []string) error {
	return nil
}

func (vers *VersionCommand) Validate(configuration *todo.Configuration) error {
	return nil
}

func (vers *VersionCommand) Help() string {
	return ""
}
