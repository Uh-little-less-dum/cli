package build

import (
	build_config "github.com/igloo1505/ulldCli/internal/build/config"
	mainBuildModel "github.com/igloo1505/ulldCli/internal/build/ui/mainmodel"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func BuildUlld(cmd *cobra.Command, dirPath string) {
	var cfg build_config.BuildConfigOpts
	cfg.Init(cmd, dirPath)
	mm := mainBuildModel.InitialMainModel(&cfg)

	tp := tea.NewProgram(mm, tea.WithAltScreen())
	if _, err := tp.Run(); err != nil {
		cobra.CheckErr(err)
	}
}
