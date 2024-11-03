package cmd

import (
	"errors"
	"ulld/cli/internal/build"

	"github.com/spf13/cobra"
)

// var (
// 	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ui.UlldBlue)).Bold(true)
// )

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new ULLD application.",
	Long:  "Builds a new ULLD application based on local configuration files and environment variables.",
	Args:  cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		env := build.UlldEnv{}
		env.Init()
		var dirPath string
		if len(args) == 0 {
			dirPath = args[0]
		}
		build.BuildUlld(env, cmd, dirPath)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// buildCmd.Root().InitDefaultHelpCmd()
}
