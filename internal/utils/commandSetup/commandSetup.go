package command_setup

import (
	cli_config "ulld/cli/internal/utils/initViper"
	"ulld/cli/internal/utils/logger"

	"github.com/spf13/cobra"
)

func InitializeCommand(loggerPrefix string, cmd *cobra.Command) func() {
	logger.InitLogger(loggerPrefix)
	return cli_config.InitViper(cmd)
}
