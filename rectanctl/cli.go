package main

import (
	"fmt"
	"os"

	"github.com/nirav24/rectangle-assessment/command"
)

func main() {
	c := command.Command()
	if err := c.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
