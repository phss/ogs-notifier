package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ogscli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ogscli v0.1")
	},
}
