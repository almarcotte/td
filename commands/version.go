package commands

import (
	"fmt"
	"github.com/gnumast/td/cli"
	"runtime"
)

const (
	VERSION = "0.0.1"
)

type VersionCommand struct{}

func (vers VersionCommand) Run(app *cli.Application) error {
	fmt.Printf("td version %s %s/%s\n", VERSION, runtime.GOOS, runtime.GOARCH)

	return nil
}

func (vers VersionCommand) Parse(app *cli.Application, args []string) error {
	return nil
}

func (vers VersionCommand) Validate(app *cli.Application) error {
	return nil
}

func (vers VersionCommand) Help() string {
	return ""
}
