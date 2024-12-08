package build_stage_utils

import (
	"path/filepath"

	fs_utils "github.com/Uh-little-less-dum/build/pkg/fs"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/spf13/viper"
)

func GetNextBuildStage() (configPath string, stage constants.BuildStage) {
	v := viper.GetViper()
	configDir := v.GetString(string(viper_keys.ConfigDir))
	if configDir != "" {
		configPath := filepath.Join(configDir, "appConfig.ulld.json")
		if fs_utils.Exists(configPath) {
			v.Set(string(viper_keys.AppConfigPath), configPath)
			return configPath, constants.ConfirmConfigLocFromEnv
		} else {
			return "", constants.ChooseWaitOrPickConfigLoc
		}
	} else {
		return "", constants.ChooseWaitOrPickConfigLoc
	}
}
