package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(peekCmd)
}

var peekCmd = &cobra.Command{
	Use:   "peek",
	Short: "Get a quick glance of a CSV file",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Peek a file")
	},
}