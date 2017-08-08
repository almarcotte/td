package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
)

type FindCommand struct {
}

func (find *FindCommand) Run(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (find *FindCommand) Validate(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (find *FindCommand) Parse(app *todo.Application, args []string) error {
	return errors.New("Not implemented")
}

func (find FindCommand) Help() string {
	return ""
}
