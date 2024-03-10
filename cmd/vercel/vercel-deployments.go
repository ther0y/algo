package vercelCmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	. "masood.dev/algo/services/vercel-client"
)

var projectID string

func init() {
	vercelCmd.AddCommand(vercelDeploymentsCmd)
	vercelDeploymentsCmd.PersistentFlags().StringVarP(&projectID, "project-id", "p", "", "Project ID")
	vercelDeploymentsCmd.PersistentFlags().BoolVarP(&asFilter, "as-filter", "f", false, "Return results as filter")
}

var vercelDeploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "deployments command",
	Long:  `deployments command`,
	Run:   runVercelDeploymentsCmd,
}

func runVercelDeploymentsCmd(cmd *cobra.Command, args []string) {
	if projectID == "" {
		cmd.Println("Project ID is required")
		return
	}

	deployments, err := GetDeploymentsByProjectID(projectID)

	if err != nil {
		cmd.Println("Error getting deployments:", err)
		return
	}

	if asFilter {
		deployments.SendAsAlfredFilter()
		return

	}

	for _, deployment := range deployments {
		printPrettyJSON(deployment)
	}
}

// printPrettyJSON prints the given response in a pretty-printed JSON format
func printPrettyJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(b))
}
