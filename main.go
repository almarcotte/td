package main

import (
	"fmt"
	"github.com/gnumast/td/cli"
	"github.com/gnumast/td/commands"
	"os"
)

func main() {
	handler := commands.NewHandler(os.Args)
	application := cli.NewApplication()

	result, err := handler.Run(application)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%s\n", result)
}
