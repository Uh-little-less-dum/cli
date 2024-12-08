package cmd_init

import (
	"os"

	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func Build(args []string, b *build_config.BuildManager) {
	v := viper.GetViper()
	useCwd := v.GetBool(string(viper_keys.UseCwd))
	var targetDir string
	if len(args) > 0 {
		targetDir = args[0]
	}
	if (targetDir != "") && (useCwd) {
		log.Fatal("Cannot provide both the `--here` flag and a positional argument. That indicates that you would like to use 2 separate paths for the same process.")
	}
	if useCwd {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		targetDir = cwd
		b.AddSkippedStage(constants.PickTargetDirStage)
		b.AddSkippedStage(constants.ConfirmCurrentDirStage)
	}
	if targetDir != "" {
		b.TargetDir = targetDir
		// b.SetTargetDir(targetDir)
	}
}
