package clone_template_app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/igloo1505/ulldCli/internal/build/constants"
	clone_template_app_manager "github.com/igloo1505/ulldCli/internal/build/stages"
	stdout_wrapper "github.com/igloo1505/ulldCli/internal/build/ui/stdoutWrapper"
	"github.com/igloo1505/ulldCli/internal/signals"
	"github.com/spf13/viper"
)

type CloneStatus int

const (
	NotStarted CloneStatus = iota
	Running
	Complete
)

type Model struct {
	outputWrapper stdout_wrapper.Model
	Id            string
	Stage         constants.BuildStage
	status        CloneStatus
}

type keymap struct {
	Quit key.Binding
}

var Keymap = keymap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case signals.SetStageMsg:
		if (m.status == NotStarted) && (msg.NewStage == m.Stage) {
			targetDir := viper.GetViper().GetString("targetDir")
			if targetDir == "" {
				log.Fatal("Attempted to build ULLD in an invalid location.")
			}
			m.beginSparseClone(targetDir)
			m.status = Running
		}
	case signals.StdOutWrapperOutputMsg:
		outputModel, cmd := m.outputWrapper.Update(msg)
		m.outputWrapper = outputModel
		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keymap.Quit):
			quitMsg := signals.SetQuittingMessage(nil)
			return m, quitMsg
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.outputWrapper.View()
}

// func (m Model) Write(p []byte) (n int, err error) {
// 	return 0, nil
// }

func (m Model) beginSparseClone(targetDir string) {
	clone_template_app_manager.CloneTemplateApp(targetDir, m.outputWrapper)
}

func NewCloneTemplateAppUIModel() Model {
	id := "clone-template-app"
	targetDir := viper.GetViper().GetString("targetDir")
	initialString := fmt.Sprintf("Cloning %s into %s...", constants.SparseCloneRepoUrl, targetDir)
	return Model{
		outputWrapper: stdout_wrapper.NewModel(initialString),
		Id:            id,
		Stage:         constants.CloneTemplateAppStage,
		status:        NotStarted,
	}
}
