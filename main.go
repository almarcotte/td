package main

import (
	"fmt"
	"github.com/gnumast/td/commands"
	"os"
)

const (
	VERSION = "0.0.1"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	if args[0] == "version" {
		version()
		os.Exit(0)
	}

	if args[0] == "config" {
		fmt.Println("Not implemented yet. Sorry!")
		os.Exit(0)
	}

	if err := commands.Handle(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		defer os.Exit(1)
	}
}

func usage() {
	version()
	fmt.Println("Usage:\n  command [arguments]\n  command help")

	fmt.Println()

	fmt.Println("Available commands:")
	fmt.Println("  a             Add a new item")
	fmt.Println("  c             Mark an item as completed")
	fmt.Println("  p             Mark an item as in progress")
	fmt.Println("  r             Revert a status to incomplete status")
	fmt.Println("  t             Add tags to an item")
	fmt.Println("  f             Find an item")
	fmt.Println("  ls            List all items")

	fmt.Println()

	fmt.Println("Global commands:")
	fmt.Println("  ga            Add a new item to the global list")
	fmt.Println("  gf            Find an item in the global list")
	fmt.Println("  gl            List all the items in the global list")

	fmt.Println()

	fmt.Println("Other commands:")
	fmt.Println("  version       Displays version information")
	fmt.Println("  config        Lists configuration options")
}

func version() {
	fmt.Printf("todo %s\n\n", VERSION)
}
