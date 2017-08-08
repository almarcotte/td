package commands

import (
	"fmt"
	"github.com/gnumast/td/cli"
	"runtime"
	"testing"
)

func TestVersionCommand_Run(t *testing.T) {
	versionCmd := VersionCommand{}
	output, stdOutBuffer, _ := cli.NewTestOutput()

	versionCmd.Run(&cli.Application{CliOutput: output})

	expected := fmt.Sprintf(VERSION_FORMAT+"\n", VERSION, runtime.GOOS, runtime.GOARCH)

	if result := stdOutBuffer.String(); result != expected {
		t.Fatalf("Expected `%s`, got `%s`", expected, result)
	}
}

func TestVersionCommand_RunThroughHandler(t *testing.T) {
	handler := NewHandler([]string{"td", "version"})
	output, stdOutBuffer, _ := cli.NewTestOutput()

	handler.Run(&cli.Application{CliOutput: output})

	expected := fmt.Sprintf(VERSION_FORMAT+"\n", VERSION, runtime.GOOS, runtime.GOARCH)

	if result := stdOutBuffer.String(); result != expected {
		t.Fatalf("Expected `%s`, got `%s`", expected, result)
	}
}
