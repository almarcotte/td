package commands

import "github.com/gnumast/td/todo"

type TagCommand struct {
	Tags []string
}

func (c *TagCommand) Validate(configuration *todo.Configuration) error {
	return nil
}

func (c *TagCommand) Parse(configuration *todo.Configuration, args []string) error {
	return nil
}

func (c *TagCommand) Run(configuration *todo.Configuration) error {
	return nil
}
