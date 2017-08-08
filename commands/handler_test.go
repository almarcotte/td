package commands

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHandler(t *testing.T) {
	args := []string{"zero", "one", "two", "three"}
	handler := NewHandler(args)

	if argsOnly := args[1:]; !reflect.DeepEqual(argsOnly, handler.Args) {
		t.Fatalf("Expected handler arguments to be %+v, got %+v", argsOnly, handler.Args)
	}
}

func TestNewHandler_NoArguments(t *testing.T) {
	args := []string{"zero"}
	handler := NewHandler(args)

	if argsOnly := args[1:]; !reflect.DeepEqual(argsOnly, handler.Args) {
		t.Fatalf("Expected handler arguments to be empty array, got %+v", handler.Args)
	}
}

func TestParseForCommand(t *testing.T) {
	handler := Handler{}
	expected := map[string]string{
		addCmd:      "*commands.AddCommand",
		completeCmd: "*commands.StatusCommand",
		progressCmd: "*commands.StatusCommand",
		revertCmd:   "*commands.StatusCommand",
		tagCmd:      "*commands.TagCommand",
		listCmd:     "*commands.ListCommand",
		versionCmd:  "*commands.VersionCommand",
		"":          "*commands.UsageCommand",
	}

	for arg, cmd := range expected {
		handler.Args = []string{arg}
		if created, _, err := handler.ParseForCommand(); fmt.Sprintf("%T", created) != cmd || err != nil {
			t.Fatalf("Expected `%s` to create %s, created %+v", arg, cmd, fmt.Sprintf("%T", created))
		}
	}
}
