package package_manager_select

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	general_select "github.com/Uh-little-less-dum/cli/internal/build/ui/generalSelect"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	tea "github.com/charmbracelet/bubbletea"
)

func NewModel() general_select.Model {
	availablePackageManagers := package_managers.GetAvailablePackageManagers()
	packageManagerNames := package_managers.GetPackageManagerTitles()

	opts := []string{}

	for k, v := range availablePackageManagers {
		if v {
			opts = append(opts, packageManagerNames[k])
		}
	}

	m := general_select.NewModel(opts, "Which package manager would you like to use?", func(packageManagerTitle string) tea.Cmd {
		return func() tea.Msg {
			var pm package_managers.PackageManagerId = package_managers.NoPackagekManagerSelected
			for x, k := range availablePackageManagers {
				if k {
					pm = x
					break
				}
			}
			build_config.GetBuildManager().SetPackageManager(pm)
			return signals.SetStageMsg{
				NewStage: build_stages.CloneTemplateAppStage,
			}
		}
	}, build_stages.SelectPackageManager)

	m.SetShowStatus(false)
	m.SetAllowFilter(false)

	return m
}
