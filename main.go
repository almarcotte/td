package main

import (
	"fmt"
	"github.com/gnumast/td/cli"
	"os"
)

func main() {
	handler := cli.NewHandler(os.Args)
	config := cli.NewConfiguration()

	if err := handler.Run(config); err != nil {
		fmt.Printf("Error! %v\n", err)
		os.Exit(1)
	}
}
