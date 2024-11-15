package mainBuildModel

import (
	buildConfig "ulld/cli/internal/build/config"
	"ulld/cli/internal/build/constants"
	keyMap "ulld/cli/internal/build/keymap"
	"ulld/cli/internal/build/ui/confirmdir"
	dirPicker "ulld/cli/internal/build/ui/dirpicker"
	"ulld/cli/internal/keymap"
	"ulld/cli/internal/signals"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	stage           constants.BuildStage
	help            help.Model
	keys            keyMap.KeyMap
	confirmDirModel confirmdir.Model
	targetDirModel  *dirPicker.DirPickerModel
	targetDir       string
	quitting        bool
}

func (m mainModel) Init() tea.Cmd {
	tea.SetWindowTitle("ULLD Build")
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case signals.SetStageMsg:
		m.stage = msg.NewStage
	case signals.SetUseSelectedDirMsg:
		if msg.UseSelectedDir {
			m.stage = constants.CloneTemplateAppStage
		} else {
			m.stage = constants.PickTargetDirStage
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
			// default:
			// 	m.confirmDirModel, cmd = m.confirmDirModel.Update(msg)
		}
	}

	switch m.stage {
	case m.confirmDirModel.Stage:
		m.confirmDirModel, cmd = m.confirmDirModel.Update(msg)
		cmds = append(cmds, cmd)
	case m.targetDirModel.Stage:
		return m.targetDirModel.Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
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
		confirmDirModel: confirmdir.NewModel("Do you want to build ULLD in your current directory?"),
		targetDir:       cfg.TargetDir,
		quitting:        false,
	}

	return &val
}
