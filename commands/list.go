package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
)

type ListCommand struct{}

func (list *ListCommand) Validate(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (list *ListCommand) Parse(app *todo.Application, args []string) error {
	return errors.New("Not implemented")
}

func (list *ListCommand) Run(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (list ListCommand) Help() string {
	return ""
}
