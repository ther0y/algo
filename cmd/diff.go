package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"masood.dev/algo/agents"
)

var RepoDir string
var DiffCmd string
var AiReview bool

func init() {
	RootCmd.AddCommand(diffCmd)

	diffCmd.PersistentFlags().StringVarP(&DiffCmd, "diff-cmd", "c", "git diff", "Diff command")
	diffCmd.PersistentFlags().StringVarP(&RepoDir, "repo-dir", "d", "", "Repo directory")
	diffCmd.PersistentFlags().BoolVarP(&AiReview, "ai-review", "a", false, "AI review")
}

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "diff command",
	Long:  `diff command`,
	Run: func(cmd *cobra.Command, args []string) {
		diffCommand := DiffCmd + " -- . ':!package-lock.json'"

		command := exec.Command("sh", "-c", diffCommand)
		command.Dir = RepoDir

		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error running diff command:", err)
			return
		}

		if AiReview {
			developerAgent := agents.NewDeveloperAgent()
			rest := developerAgent.Ask("Here is the git diff result of a merge request please review the code with all the details and suggest change requests => " + strings.TrimSpace(string(output)))

			fmt.Println(rest.Choices[0].Message.Content)
		} else {
			fmt.Println("Git diff output:")
			fmt.Println(strings.TrimSpace(string(output)))
		}
	},
}
