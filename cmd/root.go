package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

var (
	// Main context
	ctx = context.Background()

	// Main CLI title
	title = `---------------------
--- NxGenT v0.5.1 ---
---------------------`

	// Root CLI command
	rootCmd = &cobra.Command{
		Use:  "nxgent",
		Long: title,
	}
)

func init() {

	// CLI init configuration
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// CLI available commands
	rootCmd.AddCommand(
		deployCmd,
		registerCmd,
		runCmd,
		repInfoCmd,
		monitorCmd,
		//testCmd,
	)
}

func Execute() error {
	return rootCmd.Execute()
}
