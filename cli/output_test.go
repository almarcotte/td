package cli

import (
	"errors"
	"testing"
)

func TestNewOutputWithWriter(t *testing.T) {
	output, stdOutBuffer, stdErrBuffer := NewTestOutput()

	output.Write("Hello!")
	output.Error(errors.New("This is an error!"))

	if result := stdOutBuffer.String(); result != "Hello!\n" {
		t.Fatalf("Expected `Hello!`, got `%s`, stdOutBuffer = %+v", result, stdOutBuffer)
	}

	if result := stdErrBuffer.String(); result != "This is an error!\n" {
		t.Fatalf("Expected `This is an error!`, got `%s`, stdErrBuffer = %+v", result, stdErrBuffer)
	}
}
