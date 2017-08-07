package main

import (
	"fmt"
	"github.com/gnumast/td/commands"
	"os"
)

func main() {
	handler := commands.NewHandler(os.Args)
	config := commands.NewConfiguration()

	if err := handler.Run(config); err != nil {
		fmt.Printf("Error! %v\n", err)
		os.Exit(1)
	}
}
