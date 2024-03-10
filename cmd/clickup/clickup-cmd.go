package clickupcmd

import (
	"github.com/spf13/cobra"
	"masood.dev/algo/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(clickupCmd)
}

var clickupCmd = &cobra.Command{
	Use:   "clickup",
	Short: "ClickUp is a productivity platform that provides a fundamentally new way to work.",
	Long: `ClickUp is a productivity platform that provides a fundamentally new way to work.
It's a more intuitive way to work and is designed for people who want to get things done.`,
	Run: runClickUpCmd,
}

func runClickUpCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}
