package commands

import (
	"fmt"
)

type AddCommand struct {
	Due         string
	Tags        string
	Description string
}

func (add *AddCommand) Run(conf *Configuration) error {

}

func (add *AddCommand) Validate(conf *Configuration) error {

}

func (add *AddCommand) String() string {
	return fmt.Sprintf("Description: %s\nDue Date: %s\nTags: %v\n", add.Description, add.Due, add.Tags)
}
