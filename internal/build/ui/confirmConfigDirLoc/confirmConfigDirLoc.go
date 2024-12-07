package confirm_config_dir_loc

import (
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	general_confirm "github.com/Uh-little-less-dum/cli/internal/build/ui/generalConfirm"
	"github.com/Uh-little-less-dum/cli/internal/signals"
	"github.com/spf13/viper"
)

func NewModel() general_confirm.Model {
	m := general_confirm.NewModel("Is this the config that you would like to use?", viper.GetViper().GetString("appConfigPath"), signals.SendUseEnvConfigResponse, constants.ConfirmConfigLocFromEnv)
	*m.Value = true
	return m
}
