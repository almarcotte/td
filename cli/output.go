package cli

import (
	"bufio"
	"bytes"
	"os"
)

type Output struct {
	StdOut *bufio.Writer
	StdErr *bufio.Writer
}

// NewStdOutput returns an Output with stdout and stderr has its main outputs
func NewStdOutput() *Output {
	return &Output{
		StdOut: bufio.NewWriter(os.Stdout),
		StdErr: bufio.NewWriter(os.Stderr),
	}
}

// NewTestOutput returns an Output that can be used for testing. It also provides a bytes buffer for stdout and stderr
func NewTestOutput() (*Output, *bytes.Buffer, *bytes.Buffer) {
	var stdOutBuffer bytes.Buffer
	var stdErrBuffer bytes.Buffer

	output := &Output{
		StdOut: bufio.NewWriter(&stdOutBuffer),
		StdErr: bufio.NewWriter(&stdErrBuffer),
	}

	return output, &stdOutBuffer, &stdErrBuffer
}

// NewOutputWithWriter returns an Output but allows specifying the outputs for stdin and stdout
func NewOutputWithWriter(out *bufio.Writer, err *bufio.Writer) *Output {
	return &Output{
		StdOut: out,
		StdErr: err,
	}
}

// Write outputs a string to the standard output
func (output *Output) Write(line string) (len int, err error) {
	len, err = output.StdOut.WriteString(line)
	output.StdOut.Flush()

	return
}

// Error outputs an error message to the standard error output
func (output *Output) Error(error error) (len int, err error) {
	len, err = output.StdErr.WriteString(error.Error())
	output.StdErr.Flush()

	return
}
