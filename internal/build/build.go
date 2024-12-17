package build

import (
	build_main_model "github.com/Uh-little-less-dum/cli/internal/build/ui/mainmodel"
	cmd_init "github.com/Uh-little-less-dum/cli/internal/cmdInit"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func BuildUlld(cmd *cobra.Command, args []string) {
	cfg := cmd_init.Build(args)
	cfg.Init(args)
	mm := build_main_model.InitialMainModel(cfg)
	tp := tea.NewProgram(mm, tea.WithAltScreen())
	// WARN: Can refactor this to use the program provided to the build manager instead of providing it to the mainModel twice.
	mm.ApplyProgramProp(tp)
	if _, err := tp.Run(); err != nil {
		cobra.CheckErr(err)
	}
}
