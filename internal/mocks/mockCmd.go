package mocks

import (
	"os"

	"github.com/charmbracelet/log"
	command_setup "github.com/igloo1505/ulldCli/internal/utils/commandSetup"
	cli_config "github.com/igloo1505/ulldCli/internal/utils/initViper"
	"github.com/spf13/cobra"
)

func GetDirPath(args []string) string {
	var dirPath string
	if len(args) == 1 {
		dirPath = args[0]
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		dirPath = dir
	}
	return dirPath
}

var MockCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new ULLD application.",
	Long:  "Builds a new ULLD application based on local configuration files and environment variables.",
	Args:  cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(command_setup.InitializeCommand(MockCmd, cli_config.BuildCmdName, ""))
}
