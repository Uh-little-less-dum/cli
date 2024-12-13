package confirm_config_dir_loc

import (
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	general_confirm "github.com/igloo1505/ulldCli/internal/build/ui/generalConfirm"
	"github.com/spf13/viper"
)

func NewModel() general_confirm.Model {
	m := general_confirm.NewModel("Is this the config that you would like to use?", viper.GetViper().GetString("appConfigPath"), signals.SendUseEnvConfigResponse, build_stages.ConfirmConfigLocFromEnv)
	*m.Value = true
	return m
}
