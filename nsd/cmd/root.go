package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "nsd",
	Short:        "A simple cli downloader for nodejs, npm and yarn",
	Long:         "",
	SilenceUsage: true,
}

// Execute executes the root command.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
		return err
	}

	return nil
}
