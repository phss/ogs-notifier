package cli

import (
	"fmt"
	"os"

	"github.com/phss/ogs-notifier/pkg/cli/command"
)

// Execute runs the CLI commands.
func Execute() {
	if err := command.Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
