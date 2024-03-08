package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Query string
var AsFilter bool

var RootCmd = &cobra.Command{
	Use:   "algo",
	Short: "Algo is Masood's personal assistant",
	Long:  `Algo is a CLI tool for Masood's personal assistant. It is a tool to help Masood with his daily tasks.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
