package command

import (
	"github.com/nirav24/rectangle-assessment/command/execute"
	"github.com/spf13/cobra"
	"os"
)

// Command is the root command.
func Command() *cobra.Command {
	c := &cobra.Command{
		Use:           os.Args[0],
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	c.AddCommand(
		execute.Command(),
	)
	return c
}
