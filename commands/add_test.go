package commands

import (
	"testing"
)

func TestAddCommand_Help(t *testing.T) {
	addCmd := AddCommand{}

	if output := addCmd.Help(); output != "" {
		t.Fatalf("Expected `` from Help(), got %s", output)
	}
}
