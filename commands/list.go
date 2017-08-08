package commands

import (
	"errors"
	"github.com/gnumast/td/cli"
)

type ListCommand struct{}

func (list *ListCommand) Validate(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (list *ListCommand) Parse(app *cli.Application, args []string) error {
	return errors.New("Not implemented")
}

func (list *ListCommand) Run(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (list ListCommand) Help() string {
	return ""
}
