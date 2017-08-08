package commands

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHandler(t *testing.T) {
	args := []string{"one", "two", "three"}
	handler := NewHandler(args)

	if !reflect.DeepEqual(args, handler.Args) {
		t.Fatalf("Expected handler arguments to be %+v, got %+v", args, handler.Args)
	}
}

func TestParseForCommand(t *testing.T) {
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
		if created, err := ParseForCommand(arg); fmt.Sprintf("%T", created) != cmd || err != nil {
			t.Fatalf("Expected `%s` to create %s, created %+v", arg, cmd, fmt.Sprintf("%T", created))
		}
	}
}
