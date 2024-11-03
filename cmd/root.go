/*
Copyright © 2024 Anderw Mueller <aiglinski414@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	var cfgFile string
	var logFile string
	// rootCmd

	// NOT_IMPLEMENTED: Need to implement the entire config file handling.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $ULLD_ADDITIONAL_SOURCES/cliConfig.json)")

	rootCmd.PersistentFlags().StringVarP(&logFile, "logFile", "l", "", "Log output to this file. Useful for build failures and other debugging use cases.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
