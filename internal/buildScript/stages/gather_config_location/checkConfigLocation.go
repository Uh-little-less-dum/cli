package stage_gather_config_location

import (
	"path/filepath"

	"github.com/igloo1505/ulldCli/internal/build/constants"
	fs_utils "github.com/igloo1505/ulldCli/internal/utils/fs"
	"github.com/spf13/viper"
)

// Should be called just before cloning the app to ensure that a config file can be located.
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
