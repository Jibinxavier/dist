package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version  = "dev"
	commitID = "unknown"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "v",
	Long:  `Version of the project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hellogopher %s\n", Version)
	},
}
