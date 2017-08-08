package commands

import (
	"errors"
	"github.com/gnumast/td/cli"
)

type TagCommand struct {
	Tags []string
}

func (tag *TagCommand) Validate(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (tag *TagCommand) Parse(app *cli.Application, args []string) error {
	return errors.New("Not implemented")
}

func (tag *TagCommand) Run(app *cli.Application) error {
	return errors.New("Not implemented")
}

func (tag TagCommand) Help() string {
	return ""
}
