package nobitexCmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"masood.dev/algo/services/nobitex"
)

var srcCurrency string
var dstCurrency string

var nobitexStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Fetch nobitex stats",
	Long:  `Fetch nobitex stats.`,
	Run:   runNobitexStatsCmd,
}

func init() {
	// Register the command and its flags
	nobitexCmd.AddCommand(nobitexStatsCmd)
	nobitexStatsCmd.Flags().StringVar(&srcCurrency, "src-currency", "btc", "Source currency")
	nobitexStatsCmd.Flags().StringVar(&dstCurrency, "dst-currency", "rls", "Destination currency")
}

func runNobitexStatsCmd(cmd *cobra.Command, args []string) {
	resp, err := nobitex.GetMarketStats(srcCurrency, dstCurrency)

	if err != nil {
		fmt.Println("Error fetching nobitex stats:", err)
		return
	}

	printPrettyJSON(resp)
}
