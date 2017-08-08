package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
)

type StatusCommand struct {
	Flag   string
	Status todo.Status
}

func (status *StatusCommand) Validate(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (status *StatusCommand) Parse(app *todo.Application, args []string) error {
	return errors.New("Not implemented")
}

func (status *StatusCommand) Run(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (status StatusCommand) Help() string {
	return ""
}
