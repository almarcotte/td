package commands

import (
	"errors"
	"github.com/gnumast/td/cli"
	"regexp"
	"strings"
)

type AddCommand struct {
	Due         string
	Tags        []string
	Description string
}

const (
	TAG_EX  = `(![-\w]+)`
	DATE_EX = `(@today|@tomorrow|(@[\d]{2,4}-[\d]{2}-[\d]{2,4})|@next week)`
)

func (add *AddCommand) Run(app *cli.Application) error {
	return errors.New("Not implemented")
}

// Validate validates the values parsed previously and returns an error if it encounters any issues
func (add *AddCommand) Validate(app *cli.Application) error {
	if add.Description == "" {
		return errors.New("Description cannot be empty")
	}

	return nil
}

// Parse reads everything passed after the add command and tries to extract the description, due date and tags.
func (add *AddCommand) Parse(app *cli.Application, args []string) error {
	if len(args) == 0 {
		return errors.New("Missing argument for add command")
	}

	full := strings.Join(args, " ")

	// Get the tags from all the arguments and remove from what we'll consider the description
	description, tags, err := extractTags(full)
	if err != nil {
		return err
	}

	add.Tags = tags

	// Get the due date and remove it from all the arguments, leaving just the task description
	description, dueDate, err := extractDueDate(description)
	if err != nil {
		return err
	}

	add.Due = dueDate

	// Finally description is everything left after removing tags
	add.Description = strings.Trim(description, " ")

	return nil
}

func (add AddCommand) Help() string {
	return ""
}

func extractTags(full string) (description string, tags []string, err error) {
	exp, err := regexp.Compile(TAG_EX)

	if err != nil {
		return
	}

	description = full
	tagsResult := exp.FindAllString(full, -1)

	for _, tag := range tagsResult {
		tags = append(tags, strings.TrimLeft(tag, "!"))
		description = strings.Replace(description, tag, "", -1)
	}

	return
}

func extractDueDate(full string) (description string, dueDate string, err error) {
	exp, err := regexp.Compile(DATE_EX)

	if err != nil {
		return
	}

	description = full
	dateResult := exp.FindString(full)

	if dateResult != "" {
		description = strings.Replace(description, dateResult, "", 1)
	}

	dueDate = strings.TrimLeft(dateResult, "@")

	return
}
