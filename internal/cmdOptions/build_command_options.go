package cmd_options

import (
	"github.com/Uh-little-less-dum/cli/internal/cmd_option"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
)

func GetBuildCommandOptions() []cmd_option.CmdOption {
	data := []cmd_option.CmdOption{
		// Timeout flag
		cmd_option.CmdOptionInt{
			ViperKey:     viper_keys.CloneTimeout,
			FlagString:   "timeout",
			ShortHand:    "",
			DefaultValue: 180,
			UsageString:  "Sets the timeout limit in seconds for clone and installation requests.",
		},
		// Bypass location select flag
		cmd_option.CmdOptionBool{
			ViperKey:     viper_keys.UseCwd,
			FlagString:   "here",
			ShortHand:    "",
			DefaultValue: false,
			UsageString:  "Bypass the directory selection input and use the current working directory.",
		},
		// appConfig.ulld.json path flag
		cmd_option.CmdOptionString{
			ViperKey:     viper_keys.AppConfigPath,
			FlagString:   "appConfig",
			ShortHand:    "a",
			DefaultValue: "",
			UsageString:  "Bypass the file path select menu and use this path for your appConfig.ulld.json source.",
		},
	}

	return append(data, GetGlobalCommandOptions()...)
}
