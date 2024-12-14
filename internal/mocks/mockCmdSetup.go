package mocks

import (
	command_setup "github.com/Uh-little-less-dum/cli/internal/utils/commandSetup"
	cli_config "github.com/Uh-little-less-dum/cli/internal/utils/initViper"
)

func MockCommandSetup(cmdName cli_config.CommandName) {
	command_setup.InitializeCommand(MockCmd, cmdName, "")()
}
