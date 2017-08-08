package main

import (
	"github.com/gnumast/td/cli"
	"github.com/gnumast/td/commands"
	"os"
)

func main() {
	handler := commands.NewHandler(os.Args)
	application := cli.NewApplication()

	err := handler.Run(application)

	if err != nil {
		os.Exit(1)
	}
}
