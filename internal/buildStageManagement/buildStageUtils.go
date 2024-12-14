package build_stage_utils

import (
	"path/filepath"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	fs_utils "github.com/Uh-little-less-dum/build/pkg/fs"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/spf13/viper"
)

func GetNextBuildStage() (configPath string, stage build_stages.BuildStage) {
	v := viper.GetViper()
	configDir := v.GetString(string(viper_keys.ConfigDir))
	b := build_config.GetBuildManager()
	if configDir != "" {
		configPath := filepath.Join(configDir, "appConfig.ulld.json")
		if fs_utils.Exists(configPath) {
			if b.ConfigDirPath == "" {
				b.ConfigDirPath = configPath
			}
			return configPath, build_stages.ConfirmConfigLocFromEnv
		} else {
			return "", build_stages.ChooseWaitOrPickConfigLoc
		}
	} else {
		return "", build_stages.ChooseWaitOrPickConfigLoc
	}
}
