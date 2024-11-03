package mainBuildModel

import (
	"ulld/cli/internal/build/constants"
	confirmDir "ulld/cli/internal/build/ui/confirmdir"
	dirPicker "ulld/cli/internal/build/ui/dirpicker"

	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	stage           *constants.BuildStage
	confirmDirModel tea.Model
	targetDirModel  tea.Model
	targetDir       string
	Quitting        bool
}

func (m mainModel) Init() tea.Cmd {
	tea.SetWindowTitle("ULLD Build")
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	// var cmds []tea.Cmd

	// switch msg := msg.(type) {
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
	return m.targetDirModel, nil
}

func (m mainModel) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	return s
	// s = chosenView(m)
	// return mainStyle.Render("\n" + s + "\n\n")
}

func InitialMainModel(initialStage constants.BuildStage) *mainModel {

	val := mainModel{
		stage:           initialStage,
		targetDirModel:  dirPicker.InitialDirPicker(),
		confirmDirModel: confirmDir.InitialModel(),
		targetDir:       "",
		Quitting:        false,
	}

	return &val
}
