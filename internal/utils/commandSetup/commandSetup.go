package command_setup

import (
	"github.com/Uh-little-less-dum/cli/internal/cmd_option"
	cli_config "github.com/Uh-little-less-dum/cli/internal/utils/initViper"
	"github.com/Uh-little-less-dum/cli/internal/utils/logger"

	"github.com/spf13/cobra"
)

func InitializeCommand(cmd *cobra.Command, cmdOpts []cmd_option.CmdOption) func() {
	logger.InitLogger()
	return cli_config.InitViper(cmd, cmdOpts)
}
