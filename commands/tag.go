package commands

import (
	"github.com/gnumast/td/todo"
	"errors"
)

type TagCommand struct {
	Tags []string
}

func (tag *TagCommand) Validate(configuration *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (tag *TagCommand) Parse(configuration *todo.Configuration, args []string) error {
	return errors.New("Not implemented")
}

func (tag *TagCommand) Run(configuration *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (tag TagCommand) Help() string {
	return ""
}
