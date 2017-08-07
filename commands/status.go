package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
)

type StatusCommand struct {
	Flag   string
	Status todo.Status
}

func (status *StatusCommand) Validate(configuration *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (status *StatusCommand) Parse(configuration *todo.Configuration, args []string) error {
	return errors.New("Not implemented")
}

func (status *StatusCommand) Run(configuration *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (status StatusCommand) Help() string {
	return ""
}
