package commands

import (
	"fmt"
	"github.com/gnumast/td/cli"
	"runtime"
)

const (
	VERSION        = "0.0.1"
	VERSION_FORMAT = "td version %s %s/%s"
)

type VersionCommand struct{}

func (vers VersionCommand) Run(app *cli.Application) (err error) {
	app.Output.Write(fmt.Sprintf(VERSION_FORMAT, VERSION, runtime.GOOS, runtime.GOARCH))

	return
}

// Parse does nothing for this command
func (vers VersionCommand) Parse(app *cli.Application, args []string) (err error) {
	return
}

// Validate does nothing for this command
func (vers VersionCommand) Validate(app *cli.Application) (err error) {
	return
}

// Help does nothing for this command
func (vers *VersionCommand) Help(app *cli.Application) {
}
