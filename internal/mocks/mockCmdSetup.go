package mocks

import (
	command_setup "github.com/igloo1505/ulldCli/internal/utils/commandSetup"
	cli_config "github.com/igloo1505/ulldCli/internal/utils/initViper"
)

func MockCommandSetup(cmdName cli_config.CommandName) {
	command_setup.InitializeCommand(MockCmd, cmdName, "")()
}
