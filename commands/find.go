package commands

import (
	"errors"
	"github.com/gnumast/td/cli"
)

type FindCommand struct {
}

func (find *FindCommand) Run(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (find *FindCommand) Validate(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (find *FindCommand) Parse(app *cli.Application, args []string) error {
	return errors.New("Not implemented")
}

func (find FindCommand) Help() string {
	return ""
}
