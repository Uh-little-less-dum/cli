package cmd_init

import (
	"path/filepath"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func Build(args []string) *build_config.BuildManager {
	b := build_config.GetBuildManager()
	v := viper.GetViper()
	var targetDir string
	if len(args) > 0 {
		targetDir = args[0]
		absPath, err := filepath.Abs(targetDir)
		if err != nil {
			log.Fatal(err)
		}
		b.TargetDir = absPath
	}
	useCwd := v.GetBool(string(viper_keys.UseCwd))
	if (targetDir != "") && (useCwd) {
		log.Fatal("Cannot provide both the `--here` flag and a positional argument. That indicates that you would like to use 2 separate paths for the same process.")
	}
	if useCwd {
		b.AddSkippedStage(build_stages.PickTargetDirStage)
		b.AddSkippedStage(build_stages.ConfirmCurrentDirStage)
	}
	return b
}
