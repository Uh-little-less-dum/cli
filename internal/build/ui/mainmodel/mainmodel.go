package mainBuildModel

import (
	buildConfig "ulld/cli/internal/build/config"
	"ulld/cli/internal/build/constants"
	keyMap "ulld/cli/internal/build/keymap"
	confirmDir "ulld/cli/internal/build/ui/confirmdir"
	dirPicker "ulld/cli/internal/build/ui/dirpicker"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	stage           constants.BuildStage
	help            help.Model
	keys            keyMap.KeyMap
	confirmDirModel *confirmDir.ConfirmCurrentDirModel
	targetDirModel  *dirPicker.DirPickerModel
	targetDir       string
	quitting        bool
}

func (m mainModel) Init() tea.Cmd {
	tea.SetWindowTitle("ULLD Build")
	return nil
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	// var cmds []tea.Cmd

	// switch msg := msg.(type) {
	// FIX: Need to fix issue with these constants so that this switch statement works. Check the project from bashBunni when on wifi again.
	switch msg.(type) {
	case constants.ToRootModelMsg:
		m.stage = constants.ConfirmCurrentDirStage
	case constants.ConfirmDirectoryMsg:
		m.stage = constants.CloneTemplateAppStage
	}

	// switch m.stage {
	// case GetTargetDirStage:
	// 	newTargetDirModel, newCmd := m.targetDirModel.Update(msg)
	// 	tdModel, ok := newTargetDirModel.()
	// }

	// switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	// If we set a width on the help menu it can gracefully truncate
	// 	// its view as needed.
	// 	m.help.Width = msg.Width

	// case tea.KeyMsg:
	// 	switch {
	// 	case key.Matches(msg, m.keys.Up):
	// 		m.lastKey = "↑"
	// 	case key.Matches(msg, m.keys.Down):
	// 		m.lastKey = "↓"
	// 	case key.Matches(msg, m.keys.Left):
	// 		m.lastKey = "←"
	// 	case key.Matches(msg, m.keys.Right):
	// 		m.lastKey = "→"
	// 	case key.Matches(msg, m.keys.Help):
	// 		m.help.ShowAll = !m.help.ShowAll
	// 	case key.Matches(msg, m.keys.Quit):
	// 		m.quitting = true
	// 		return m, tea.Quit
	// 	}
	// }

	switch m.stage {
	case m.confirmDirModel.Stage:
		return m.confirmDirModel.Update(msg)
	case m.targetDirModel.Stage:
		return m.targetDirModel.Update(msg)
	}

	return m.targetDirModel, nil
}

func (m *mainModel) View() string {
	var s string
	if m.quitting {
		return "\n  No worries.\n\n"
	}
	switch m.stage {
	case constants.ConfirmCurrentDirStage:
		return m.confirmDirModel.View()
	case constants.PickTargetDirStage:
		return m.targetDirModel.View()
	}
	return s
	// s = chosenView(m)
	// return mainStyle.Render("\n" + s + "\n\n")
}

func InitialMainModel(cfg *buildConfig.BuildConfigOpts) *mainModel {

	val := mainModel{
		stage:           cfg.InitialStage,
		help:            help.New(),
		keys:            keyMap.DefaultKeymap,
		targetDirModel:  dirPicker.InitialDirPicker(),
		confirmDirModel: confirmDir.InitialModel("Do you want to build ULLD in your current directory?", ""),
		targetDir:       "",
		quitting:        false,
	}

	return &val
}
