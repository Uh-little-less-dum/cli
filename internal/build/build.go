package build

import (
	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	mainBuildModel "github.com/Uh-little-less-dum/cli/internal/build/ui/build_main_model"
	cmd_init "github.com/Uh-little-less-dum/cli/internal/cmdInit"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func BuildUlld(cmd *cobra.Command, args []string) {
	cfg := build_config.GetBuildManager()
	cmd_init.Build(args, cfg)
	cfg.Init()
	mm := mainBuildModel.InitialMainModel(cfg)

	tp := tea.NewProgram(mm, tea.WithAltScreen())
	if _, err := tp.Run(); err != nil {
		cobra.CheckErr(err)
	}
}
