package clickupcmd

import (
	"github.com/spf13/cobra"
	clickupservice "masood.dev/algo/services/clickup"
)

func init() {
	clickupCmd.AddCommand(clickupTeamsCmd)
}

var clickupTeamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Get all teams",
	Long:  `Get all teams from ClickUp`,
	Run:   runClickUpTeamsCmd,
}

func runClickUpTeamsCmd(cmd *cobra.Command, args []string) {
	clickupservice.GetTeams()
}
