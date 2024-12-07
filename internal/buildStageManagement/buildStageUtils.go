package build_stage_utils

import (
	"path/filepath"

	fs_utils "github.com/Uh-little-less-dum/build/pkg/fs"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	"github.com/spf13/viper"
)

func GetNextBuildStage() (configPath string, stage constants.BuildStage) {
	v := viper.GetViper()
	configDir := v.GetString("configDir")
	if configDir != "" {
		configPath := filepath.Join(configDir, "appConfig.ulld.json")
		if fs_utils.Exists(configPath) {
			v.Set("appConfigPath", configPath)
			return configPath, constants.ConfirmConfigLocFromEnv
		} else {
			return "", constants.ChooseWaitOrPickConfigLoc
		}
	} else {
		return "", constants.ChooseWaitOrPickConfigLoc
	}
}
