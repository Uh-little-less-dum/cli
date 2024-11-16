package command_setup

import (
	cli_config "github.com/igloo1505/ulldCli/internal/utils/initViper"
	"github.com/igloo1505/ulldCli/internal/utils/logger"

	"github.com/spf13/cobra"
)

func InitializeCommand(cmd *cobra.Command, commandName cli_config.CommandName, loggerPrefix string) func() {
	logger.InitLogger(loggerPrefix)
	return cli_config.InitViper(cmd, commandName)
}
