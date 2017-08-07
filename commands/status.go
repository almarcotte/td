package commands

import "github.com/gnumast/td/todo"

type StatusCommand struct {
	Flag   string
	Status todo.Status
}

func (c *StatusCommand) Validate(configuration *todo.Configuration) error {
	return nil
}

func (c *StatusCommand) Parse(configuration *todo.Configuration, args []string) error {
	return nil
}

func (c *StatusCommand) Run(configuration *todo.Configuration) error {
	return nil
}
