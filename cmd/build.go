package cmd

import (
	"github.com/Uh-little-less-dum/cli/internal/build"
	cmd_options "github.com/Uh-little-less-dum/cli/internal/cmdOptions"
	command_setup "github.com/Uh-little-less-dum/cli/internal/utils/commandSetup"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new ULLD application.",
	Long:  "Builds a new ULLD application based on local configuration files and environment variables.",
	Args:  cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		build.BuildUlld(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)
	command_setup.InitializeCommand(buildCmd, cmd_options.GetBuildCommandOptions())()
}
