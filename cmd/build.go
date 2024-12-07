package cmd

import (
	"github.com/Uh-little-less-dum/cli/internal/build"
	command_setup "github.com/Uh-little-less-dum/cli/internal/utils/commandSetup"
	cli_config "github.com/Uh-little-less-dum/cli/internal/utils/initViper"

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

// RESUME: Implement newly created cli option models through a list that is uniique to each command and a single function that simply accepts the list of flags and the cmd that initializes each command.
func init() {
	RootCmd.AddCommand(buildCmd)
	command_setup.InitializeCommand(buildCmd, cli_config.BuildCmdName, "")()

}
