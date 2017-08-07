package commands

import (
	"github.com/gnumast/td/todo"
	"strings"
	"testing"
)

func TestAddCommand_Help(t *testing.T) {
	addCmd := AddCommand{}

	if output := addCmd.Help(); output != "" {
		t.Fatalf("Expected `` from Help(), got %s", output)
	}
}

func TestAddCommand_ParseEmpty(t *testing.T) {
	addCmd := AddCommand{}
	conf := &todo.Configuration{}

	err := addCmd.Parse(conf, []string{})

	if err == nil {
		t.Fatal("Expected error when parsing empty string, got nil")
	}
}

func TestAddCommand_ParseDescriptionOnly(t *testing.T) {
	addCmd := AddCommand{}
	conf := &todo.Configuration{}
	description := "write tests for td"

	err := addCmd.Parse(conf, strings.Split(description, " "))

	if err != nil {
		t.Fatalf("Unexpected error when parsing: %s", err.Error())
	}

	if addCmd.Description != description {
		t.Fatalf("Unexpected parsed description. Expected %s, got %s\n", description, addCmd.Description)
	}
}

func TestAddCommand_ParseDescriptionWithTags(t *testing.T) {
	addCmd := AddCommand{}
	conf := &todo.Configuration{}
	description := "write tests for td"
	full := description + " !golang !tdd"

	err := addCmd.Parse(conf, strings.Split(full, " "))

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
