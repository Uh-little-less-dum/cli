package select_package_manager

import (
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	general_select "github.com/Uh-little-less-dum/cli/internal/build/ui/generalSelect"
	"github.com/Uh-little-less-dum/cli/internal/signals"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

func sendSetPackageManager(acceptedTitle string) tea.Cmd {
	viper.GetViper().Set("packageManager", acceptedTitle)
	return func() tea.Msg {
		return signals.SetStageMsg{
			NewStage: constants.InstallTemplateAppDeps,
		}
	}
}

// Model sets the packageManager field in viper and moves on to the InstallTemplateAppDeps stage.
func NewModel() general_select.Model {
	// pnpm := general_select.Item{}
	// pnpm.SetTitle("pnpm")

	// npm := general_select.Item{}
	// npm.SetTitle("npm")
	// yarn := general_select.Item{}
	// yarn.SetTitle("yarn")

	// opts := []general_select.Item{
	// 	pnpm,
	// 	npm,
	// 	yarn,
	// }

	opts := []string{"pnpm", "npm", "yarn"}

	m := general_select.NewModel(opts, "Which package manager would you like to use?", sendSetPackageManager, constants.ChooseWaitOrPickConfigLoc)
	return m
}
