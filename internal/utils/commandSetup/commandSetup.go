package command_setup

import (
	cli_config "github.com/igloo1505/ulldCli/internal/utils/initViper"
	"github.com/igloo1505/ulldCli/internal/utils/logger"

	"github.com/spf13/cobra"
)

func InitializeCommand(loggerPrefix string, cmd *cobra.Command) func() {
	logger.InitLogger(loggerPrefix)
	return cli_config.InitViper(cmd)
}
