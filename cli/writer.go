package cli

import (
	"bufio"
	"bytes"
	"os"
)

type Writer struct {
	StdOut *bufio.Writer
	StdErr *bufio.Writer
}

// NewStdOutput returns an CliOutput with stdout and stderr has its main outputs
func NewStdOutput() *Writer {
	return &Writer{
		StdOut: bufio.NewWriter(os.Stdout),
		StdErr: bufio.NewWriter(os.Stderr),
	}
}

// NewTestOutput returns an CliOutput that can be used for testing. It also provides a bytes buffer for stdout and stderr
func NewTestOutput() (*Writer, *bytes.Buffer, *bytes.Buffer) {
	var stdOutBuffer bytes.Buffer
	var stdErrBuffer bytes.Buffer

	output := &Writer{
		StdOut: bufio.NewWriter(&stdOutBuffer),
		StdErr: bufio.NewWriter(&stdErrBuffer),
	}

	return output, &stdOutBuffer, &stdErrBuffer
}

// Write outputs a string to the standard output
func (output *Writer) Write(line string) (len int, err error) {
	len, err = output.StdOut.WriteString(line + "\n")
	output.StdOut.Flush()

	return
}

// Error outputs an error message to the standard error output
func (output *Writer) Error(error error) (len int, err error) {
	len, err = output.StdErr.WriteString(error.Error() + "\n")
	output.StdErr.Flush()

	return
}
