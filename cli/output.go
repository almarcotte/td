package cli

import (
	"bufio"
	"fmt"
	"os"
)

type Output struct {
	*bufio.Writer
}

func NewStdOutput() *Output {
	return &Output{bufio.NewWriter(os.Stdout)}
}

func NewOutputWithWriter(writer *bufio.Writer) *Output {
	return &Output{writer}
}

func (output *Output) WriteLine(line string) (int, error) {
	return output.WriteString(fmt.Sprintf("%s\n", line))
}
