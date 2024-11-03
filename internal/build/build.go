package build

import (
	buildConfig "ulld/cli/internal/build/config"
	mainBuildModel "ulld/cli/internal/build/ui/mainmodel"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func BuildUlld(e UlldEnv, cmd *cobra.Command, dirPath string) {
	var cfg buildConfig.BuildConfigOpts
	cfg.Init(cmd, dirPath)
	mm := mainBuildModel.InitialMainModel(&cfg)

	tp := tea.NewProgram(mm, tea.WithAltScreen())
	if _, err := tp.Run(); err != nil {
		cobra.CheckErr(err)
	}
}
