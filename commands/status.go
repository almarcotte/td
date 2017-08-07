package commands

import "github.com/gnumast/td/todo"

type StatusCommand struct {
	Flag   string
	Status todo.Status
}

func (c *StatusCommand) Validate(configuration *Configuration) error {
	return nil
}

func (c *StatusCommand) Parse(configuration *Configuration, args []string) error {
	return nil
}

func (c *StatusCommand) Run(configuration *Configuration) error {
	return nil
}
