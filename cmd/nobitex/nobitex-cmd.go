package nobitexCmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"masood.dev/algo/cmd"
	"masood.dev/algo/services/nobitex"
)

func init() {
	// Add the command to the root command
	cmd.RootCmd.AddCommand(nobitexCmd)
}

var nobitexCmd = &cobra.Command{
	Use:   "nobitex",
	Short: "nobitex command",
	Long:  `nobitex command`,
	Run: func(cmd *cobra.Command, args []string) {
		// write your code here
		resp, err := nobitex.GetWalletsByCurrency("btc")

		if err != nil {
			fmt.Println(err)
		}

		// JSON pretty print
		b, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	},
}
