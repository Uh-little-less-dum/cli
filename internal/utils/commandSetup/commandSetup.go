package command_setup

import (
	"github.com/Uh-little-less-dum/cli/internal/cmd_option"
	cli_config "github.com/Uh-little-less-dum/cli/internal/utils/initViper"
	"github.com/Uh-little-less-dum/cli/internal/utils/logger"

	"github.com/spf13/cobra"
)

func InitializeCommand(cmd *cobra.Command, commandName cli_config.CommandName, loggerPrefix string) func() {
	logger.InitLogger(loggerPrefix)
	return cli_config.InitViper(cmd, commandName)
}

// TODO: Figure out how to implement 'any' type in Go, and use that here as to not pass the problem upstream when the value can be logically handled here.
func ApplyCommandOptions[T string | int | float32 | bool](cmd *cobra.Command, opts []cmd_option.CmdOption[T]) {
	for _, o := range opts {
		o.Init(cmd)
	}
}
