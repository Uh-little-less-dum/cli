/*
Copyright © 2024 Anderw Mueller <aiglinski414@gmail.com>
*/
package cmd

import (
	"os"
	command_setup "ulld/cli/internal/utils/commandSetup"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dum",
	Short: "Uh Little Less Dum",
	Long: `A cli utility for the ULLD note taking and research framework.

This cli provides the primary ULLD build script and related commands, and will in the future, provide further functionality related to the native ULLD api.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// This should only be created in the rootCmd
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(command_setup.InitializeCommand("", RootCmd))
}
