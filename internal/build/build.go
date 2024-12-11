package build

import (
	mainBuildModel "github.com/Uh-little-less-dum/cli/internal/build/ui/build_main_model"
	cmd_init "github.com/Uh-little-less-dum/cli/internal/cmdInit"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func BuildUlld(cmd *cobra.Command, args []string) {
	cfg := cmd_init.Build(args)
	cfg.Init(args)
	mm := mainBuildModel.InitialMainModel(cfg)
	tp := tea.NewProgram(mm, tea.WithAltScreen())
	mm.ApplyProgramProp(tp)
	if _, err := tp.Run(); err != nil {
		cobra.CheckErr(err)
	}
}
