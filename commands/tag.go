package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
)

type TagCommand struct {
	Tags []string
}

func (tag *TagCommand) Validate(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (tag *TagCommand) Parse(app *todo.Application, args []string) error {
	return errors.New("Not implemented")
}

func (tag *TagCommand) Run(app *todo.Application) error {
	return errors.New("Not implemented")
}

func (tag TagCommand) Help() string {
	return ""
}
