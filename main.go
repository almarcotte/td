package main

import (
	"fmt"
	"github.com/gnumast/td/commands"
	"github.com/gnumast/td/todo"
	"os"
)

func main() {
	handler := commands.NewHandler(os.Args)
	config := todo.NewConfiguration()

	if err := handler.Run(config); err != nil {
		fmt.Printf("Error! %v\n", err)
		os.Exit(1)
	}
}
