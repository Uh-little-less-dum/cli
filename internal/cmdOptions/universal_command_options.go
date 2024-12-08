package cmd_options

import (
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/Uh-little-less-dum/cli/internal/cmd_option"
)

func GetGlobalCommandOptions() []cmd_option.CmdOption {
	return []cmd_option.CmdOption{
		// Logfile path flag
		cmd_option.CmdOptionString{
			ViperKey:     viper_keys.LogFilePath,
			FlagString:   "logFile",
			ShortHand:    "",
			DefaultValue: "",
			UsageString:  "Log output to this file. Useful for build failures and other local development.",
		},

		// Loglevel flag
		cmd_option.CmdOptionString{
			ViperKey:     viper_keys.LogLevel,
			FlagString:   "logLevel",
			ShortHand:    "",
			DefaultValue: "info",
			UsageString:  "Log level",
			EnvVar:       "ULLD_LOG_LEVEL",
		},
	}
}