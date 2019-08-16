package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ogscli",
	Short: "ogscli is a command line utility to interact with OGS",
	Long:  "A command line utility to interact with the Online Go server (https://online-go.com/)",
}

// Execute runs the CLI commands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
