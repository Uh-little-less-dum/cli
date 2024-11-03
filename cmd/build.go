package cmd

import (
	"github.com/spf13/cobra"
	"ulld/cli/internal/build"
)

// var (
// 	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ui.UlldBlue)).Bold(true)
// )

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new ULLD application.",
	Long:  "Builds a new ULLD application based on local configuration files and environment variables.",

	Run: func(cmd *cobra.Command, args []string) {
		env := build.UlldEnv{}
		env.Init()
		build.BuildUlld(env)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// buildCmd.Root().InitDefaultHelpCmd()
}
