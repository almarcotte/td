package commands

import (
	"errors"
	"github.com/gnumast/td/todo"
	"regexp"
	"strings"
)

type AddCommand struct {
	Due         string
	Tags        []string
	Description string
}

const (
	TAG_EX = `(![-\w]+)`
	//DATE_EX = `(@today|@tomorrow|(@[\d]{2,4}-[\d]{2}-[\d]{2,4})|@next week)`
)

func (add *AddCommand) Run(conf *todo.Configuration) error {
	return errors.New("Not implemented")
}

func (add *AddCommand) Validate(conf *todo.Configuration) error {
	return errors.New("Not implemented")
}

// Parse reads everything passed after the add command and tries to extract the description, due date and tags.
func (add *AddCommand) Parse(conf *todo.Configuration, args []string) error {
	if len(args) == 0 {
		return errors.New("Missing argument for add command")
	}

	full := strings.Join(args, " ")

	exp, err := regexp.Compile(TAG_EX)

	if err != nil {
		return errors.New("Couldn't compile tag regex, this is likely a bug. Sorry")
	}

	tagsResult := exp.FindAllString(full, -1)

	description := full
	tags := []string{}

	for _, tag := range tagsResult {
		description = strings.Replace(description, tag, "", -1)
		tags = append(tags, strings.TrimLeft(tag, "!"))
	}

	add.Description = strings.Trim(description, " ")
	add.Tags = tags

	return nil
}

func (add AddCommand) Help() string {
	return ""
}
