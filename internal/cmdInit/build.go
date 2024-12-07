package cmd_init

import (
	"os"

	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func Build(args []string, b *build_config.BuildConfigOpts) {
	v := viper.GetViper()
	useCwd := v.GetBool("useCwd")
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
		b.SetTargetDir(targetDir)
	}
}
