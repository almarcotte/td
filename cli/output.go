package cli

import (
	"bufio"
	"fmt"
	"os"
)

type Output struct {
	*bufio.Writer
}

func NewStdOuput() *Output {
	return &Output{bufio.NewWriter(os.Stdout)}
}

func (output *Output) WriteLine(line string) (int, error) {
	return output.WriteString(fmt.Sprintf("%s\n", line))
}
