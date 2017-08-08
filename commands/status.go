package commands

import (
	"errors"
	"github.com/gnumast/td/cli"
	"github.com/gnumast/td/todo"
)

type StatusCommand struct {
	Flag   string
	Status todo.Status
}

func (status *StatusCommand) Validate(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (status *StatusCommand) Parse(app *cli.Application, args []string) error {
	return errors.New("Not implemented")
}

func (status *StatusCommand) Run(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (status StatusCommand) Help() string {
	return ""
}
