package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
)

type FindCommand struct {
}

func (find *FindCommand) Run(conf *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (find *FindCommand) Validate(conf *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (find *FindCommand) Parse(conf *todo.Configuration, args []string) error {
	return errors.New("Not implemented")
}

func (find FindCommand) Help() string {
	return ""
}
