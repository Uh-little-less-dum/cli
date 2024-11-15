package cmd

import (
	"os"
	"ulld/cli/internal/build"
	cli_config "ulld/cli/internal/utils/initViper"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new ULLD application.",
	Long:  "Builds a new ULLD application based on local configuration files and environment variables.",
	Args:  cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		dirPath := GetDirPath(args)
		if dirPath != "" {
			viper.GetViper().Set("targetDir", dirPath)
		}
		log.Debugf("Building ULLD in %s", dirPath)
		build.BuildUlld(cmd, dirPath)
	},
}

func init() {
	cobra.OnInitialize(cli_config.InitViper(buildCmd.Root()))
	RootCmd.AddCommand(buildCmd)
}
