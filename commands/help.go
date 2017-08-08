package commands

import (
	"errors"
	"github.com/gnumast/td/cli"
)

type HelpCommand struct {
	Tags []string
}

func (help *HelpCommand) Validate(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (help *HelpCommand) Parse(app *cli.Application, args []string) error {
	return errors.New("Not implemented")
}

func (help *HelpCommand) Run(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (help *HelpCommand) Help(app *cli.Application) {
}
