package commands

import "github.com/gnumast/td/todo"

type ListCommand struct{}

func (c *ListCommand) Validate(configuration *todo.Configuration) error {
	return nil
}

func (c *ListCommand) Parse(configuration *todo.Configuration, args []string) error {
	return nil
}

func (c *ListCommand) Run(configuration *todo.Configuration) error {
	return nil
}
