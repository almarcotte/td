package commands

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin/json"
	"github.com/gnumast/td/cli"
	"github.com/gnumast/td/todo"
	"strings"
	"testing"
)

func TestAddCommand_Help(t *testing.T) {
	addCmd := AddCommand{}

	// Create a fake Application with bytes buffer as output
	output, stdOutBuffer, _ := cli.NewTestOutput()
	app := &cli.Application{CliOutput: output}

	addCmd.Help(app)

	if output := stdOutBuffer.String(); output != "" {
		t.Fatalf("Expected help to return ``, got %s", output)
	}
}

func TestAddCommand_ParseEmpty(t *testing.T) {
	addCmd := AddCommand{}
	app := &cli.Application{}

	err := addCmd.Parse(app, []string{})

	if err == nil {
		t.Fatal("Expected error when parsing empty string, got nil")
	}
}

func TestAddCommand_ParseDescriptionOnly(t *testing.T) {
	addCmd := AddCommand{}
	app := &cli.Application{}
	description := "write tests for td"

	err := addCmd.Parse(app, strings.Split(description, " "))

	if err != nil {
		t.Fatalf("Unexpected error when parsing: %s", err.Error())
	}

	if addCmd.Description != description {
		t.Fatalf("Unexpected parsed description. Expected %s, got %s\n", description, addCmd.Description)
	}
}

func TestAddCommand_ParseDescriptionWithTags(t *testing.T) {
	addCmd := AddCommand{}
	app := &cli.Application{}
	description := "write tests for td"
	full := description + " !golang !tdd"

	err := addCmd.Parse(app, strings.Split(full, " "))

	if err != nil {
		t.Fatalf("Unexpected error when parsing: %s", err.Error())
	}

	if addCmd.Description != description {
		t.Fatalf("Expected description to be `%s`, got `%s`", description, addCmd.Description)
	}

	if count := len(addCmd.Tags); count != 2 {
		t.Fatalf("Expected 2 tags, got %d", count)
	}

	if parsedTags := strings.Join(addCmd.Tags, ","); parsedTags != "golang,tdd" {
		t.Fatalf("Unexpected tags parsed, expected `%s`, got `%s`", "golang,tdd", parsedTags)
	}
}

func TestAddCommand_ParseMoreDateFormats(t *testing.T) {
	addCmd := AddCommand{}
	checks := map[string]string{
		"write tests for td @tomorrow":    "tomorrow",
		"buy milk @today":                 "today",
		"do groceries @next week":         "next week",
		"buy christmas gifts @23-12-2017": "23-12-2017",
		"buy christmas gifts @2017-12-23": "2017-12-23",
	}

	for full, date := range checks {
		addCmd.Parse(&cli.Application{}, strings.Split(full, " "))
		if addCmd.Due != date {
			t.Fatalf("Expected Due Date to be `%s`, got `%s`", date, addCmd.Due)
		}

		if strings.Contains(addCmd.Description, date) {
			t.Fatalf("Unexpected date found in description: `%s` shouldn't contain `%s`", addCmd.Description, date)
		}
	}
}

func TestAddCommand_ValidateEmptyDescriptionIsInvalid(t *testing.T) {
	addCmd := &AddCommand{}

	if err := addCmd.Validate(&cli.Application{}); err == nil {
		t.Fatal("Expected empty description to return an error, got nil")
	}
}

func TestAddCommand_Run(t *testing.T) {
	t.SkipNow() // Skip for now

	// Run() should convert an AddCommand to an Item and write it to disk
	addCmd := &AddCommand{Description: "Buy some milk"}
	item := todo.Item{Id: 1, Description: "Buy some milk", Status: todo.Incomplete} // This is what is expected

	// The idea is to decouple writing to disk so it's testable, still trying to figure out best way to do this
	var fileBuffer bytes.Buffer
	application := &cli.Application{
		CurrentFile: bufio.NewReadWriter(bufio.NewReader(&fileBuffer), bufio.NewWriter(&fileBuffer)),
	}

	addCmd.Run(application)

	jsonContent, _ := json.Marshal(item)

	// Expect the fileBuffer to have the json representation of the task
	if fileContent, expected := fileBuffer.String(), string(jsonContent); fileContent != expected {
		t.Fatalf("Expected %s, got %s", expected, fileContent)
	}
}
