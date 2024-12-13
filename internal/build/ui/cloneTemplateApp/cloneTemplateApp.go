package clone_template_app

import (
	"fmt"

	build_constants "github.com/Uh-little-less-dum/build/pkg/buildConstants"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	run_status "github.com/Uh-little-less-dum/go-utils/pkg/constants/runStatus"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	stdout_wrapper "github.com/igloo1505/ulldCli/internal/build/ui/stdoutWrapper"
	stage_clone_template_app "github.com/igloo1505/ulldCli/internal/buildScript/stages/stage_clone_template_app/createTemplateApp/clone"
	"github.com/spf13/viper"
)

type Model struct {
	outputWrapper stdout_wrapper.Model
	Id            string
	Stage         build_stages.BuildStage
	status        run_status.RunStatus
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
		if (run_status.HasNotRun(m.status)) && (msg.NewStage == m.Stage) {
			targetDir := viper.GetViper().GetString("targetDir")
			if targetDir == "" {
				log.Fatal("Attempted to build ULLD in an invalid location.")
			}
			m.beginSparseClone(targetDir)
			m.status = run_status.Running
		}
	case signals.StdOutWrapperOutputMsg:
		log.Fatal("Remove this if this is never reached.")
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

func (m Model) beginSparseClone(targetDir string) {
	stage_clone_template_app.Run(targetDir, m.outputWrapper)
}

func NewCloneTemplateAppUIModel() Model {
	id := "clone-template-app"
	targetDir := viper.GetViper().GetString("targetDir")
	initialString := fmt.Sprintf("Cloning %s into %s...", build_constants.SparseCloneRepoUrl, targetDir)
	return Model{
		outputWrapper: stdout_wrapper.NewModel(initialString),
		Id:            id,
		Stage:         build_stages.CloneTemplateAppStage,
		status:        run_status.NotStarted,
	}
}
