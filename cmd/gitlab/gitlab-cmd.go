package gitlabCmd

import (
	"github.com/spf13/cobra"
	"masood.dev/algo/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(gitlabCmd)
}

var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "gitlab command",
	Long:  `gitlab command`,
	Run:   runGitlabCmd,
}

func runGitlabCmd(cmd *cobra.Command, args []string) {
	cmd.Help()
}
