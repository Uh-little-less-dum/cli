package stage_gather_config_location

import (
	"path/filepath"

	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	fs_utils "github.com/igloo1505/ulldCli/internal/utils/fs"
	"github.com/spf13/viper"
)

// Should be called just before cloning the app to ensure that a config file can be located.
func GetNextBuildStage() (configPath string, stage build_stages.BuildStage) {
	v := viper.GetViper()
	configDir := v.GetString("configDir")
	if configDir != "" {
		configPath := filepath.Join(configDir, "appConfig.ulld.json")
		if fs_utils.Exists(configPath) {
			v.Set("appConfigPath", configPath)
			return configPath, build_stages.ConfirmConfigLocFromEnv
		} else {
			return "", build_stages.ChooseWaitOrPickConfigLoc
		}
	} else {
		return "", build_stages.ChooseWaitOrPickConfigLoc
	}
}
