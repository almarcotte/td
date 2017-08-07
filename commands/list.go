package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
)

type ListCommand struct{}

func (list *ListCommand) Validate(configuration *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (list *ListCommand) Parse(configuration *todo.Configuration, args []string) error {
	return errors.New("Not implemented")
}

func (list *ListCommand) Run(configuration *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (list *ListCommand) Help() string {
	return ""
}
