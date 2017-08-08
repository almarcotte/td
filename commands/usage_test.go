package commands

import (
	"github.com/gnumast/td/cli"
	"testing"
)

func TestUsageCommand_Run(t *testing.T) {
	usageCmd := UsageCommand{}
	output, stdOutBuffer, _ := cli.NewTestOutput()

	usageCmd.Run(&cli.Application{Output: output})

	expected := USAGE_OUTPUT + "\n"

	if result := stdOutBuffer.String(); result != expected {
		t.Fatalf("Expected `%s`, got `%s`", expected, result)
	}
}

func TestUsageCommand_RunThroughHandler(t *testing.T) {
	handlers := []*Handler{
		NewHandler([]string{"td"}),
		NewHandler([]string{"td", "not", "a", "real", "cmd"}),
	}

	expected := USAGE_OUTPUT + "\n"

	for _, handler := range handlers {
		output, stdOutBuffer, _ := cli.NewTestOutput()
		handler.Run(&cli.Application{Output: output})

		if result := stdOutBuffer.String(); result != expected {
			t.Fatalf("Expected `%s`, got `%s`", expected, result)
		}
	}
}
