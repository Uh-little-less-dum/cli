package build_stage_utils

import (
	"path/filepath"

	fs_utils "github.com/Uh-little-less-dum/build/pkg/fs"
	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/spf13/viper"
)

func GetNextBuildStage() (configPath string, stage constants.BuildStage) {
	v := viper.GetViper()
	configDir := v.GetString(string(viper_keys.ConfigDir))
	b := build_config.GetBuildManager()
	if configDir != "" {
		configPath := filepath.Join(configDir, "appConfig.ulld.json")
		if fs_utils.Exists(configPath) {
			if b.ConfigDirPath == "" {
				b.ConfigDirPath = configPath
			}
			return configPath, constants.ConfirmConfigLocFromEnv
		} else {
			return "", constants.ChooseWaitOrPickConfigLoc
		}
	} else {
		return "", constants.ChooseWaitOrPickConfigLoc
	}
}
