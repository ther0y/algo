package nobitexCmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"masood.dev/algo/services/nobitex"
)

// Global variable to hold currencies from command flags
var currencies []string

func init() {
	// Register the command and its flags
	nobitexCmd.AddCommand(walletsCmd)
	walletsCmd.Flags().StringSliceVar(&currencies, "currencies", nil, "Currencies to get wallets")
}

// walletsCmd represents the wallets command
var walletsCmd = &cobra.Command{
	Use:   "wallets",
	Short: "Fetch wallet information",
	Long:  `Fetch wallet information for specified currencies.`,
	Run:   runWalletsCmd,
}

// runWalletsCmd executes the logic for the wallets command
func runWalletsCmd(cmd *cobra.Command, args []string) {
	var resp interface{}
	var err error

	switch len(currencies) {
	case 0:
		resp, err = nobitex.GetWallets()
	case 1:
		resp, err = nobitex.GetWalletsByCurrency(currencies[0])
	default:
		resp, err = nobitex.GetWalletsByCurrencies(currencies)
	}

	if err != nil {
		fmt.Println("Error fetching wallets:", err)
		return
	}

	printPrettyJSON(resp)
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
