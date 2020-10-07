package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nsd",
	Short: "A simple cli downloader for nodejs, npm and yarn",
	Long:  "",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
