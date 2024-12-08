package confirm_app_config_loc

import (
	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	general_confirm "github.com/Uh-little-less-dum/cli/internal/build/ui/generalConfirm"
	"github.com/Uh-little-less-dum/cli/internal/signals"
)

func NewModel(buildConfig *build_config.BuildManager) general_confirm.Model {
	m := general_confirm.NewModel("Is this the config file that you would like to use?", buildConfig.AppConfigPath, signals.SendUseEnvConfigResponse, constants.ConfirmConfigLocFromEnv)
	*m.Value = true
	return m
}
