package main

import (
	"fmt"
	"github.com/gnumast/td/commands"
	"github.com/gnumast/td/todo"
	"os"
)

func main() {
	handler := commands.NewHandler(os.Args)
	config := todo.NewApplication()

	result, err := handler.Run(config)

	if err != nil {
		fmt.Printf("Error! %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", result)
}
