package commands

import "github.com/gnumast/td/todo"

type AddCommand struct {
	Due         string
	Tags        string
	Description string
}

func (add *AddCommand) Run(conf *todo.Configuration) error {
	return nil
}

func (add *AddCommand) Validate(conf *todo.Configuration) error {
	return nil
}

// Parse reads everything passed after the add command and tries to extract the description, due date and tags.
func (add *AddCommand) Parse(conf *todo.Configuration, args []string) error {
	return nil
}
