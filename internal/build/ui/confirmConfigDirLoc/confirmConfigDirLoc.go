package confirm_app_config_loc

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	general_confirm "github.com/Uh-little-less-dum/cli/internal/build/ui/generalConfirm"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
)

func NewModel(buildConfig *build_config.BuildManager) general_confirm.Model {
	m := general_confirm.NewModel("Is this the config file that you would like to use?", buildConfig.AppConfigPath, signals.SendUseEnvConfigResponse, build_stages.ConfirmConfigLocFromEnv)
	*m.Value = true
	return m
}
