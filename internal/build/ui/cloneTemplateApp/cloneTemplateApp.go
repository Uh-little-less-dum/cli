package clone_template_app

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	stage_clone_template_app "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/stage_clone_template_app/createTemplateApp/clone"
	general_spinner "github.com/Uh-little-less-dum/cli/internal/build/ui/generalSpinner"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type CloneStatus int

type responseMsg bool

const (
	NotStarted CloneStatus = iota
	Running
	Complete
)

type Model struct {
	Id      string
	Stage   build_stages.BuildStage
	status  CloneStatus
	spinner general_spinner.Model
	sub     chan bool
	Program *tea.Program
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

func (m Model) waitForActivity(sub chan bool) tea.Cmd {
	return func() tea.Msg {
		return responseMsg(<-sub)
	}
}

func (m Model) cloneFinishedMsg() tea.Cmd {
	return signals.SetStage(build_stages.PreConflictResolveBuild)
}

// RESUME: Come back here and handle the responseMsg. This fucking thing worked once randomly and broke again.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	s, cmd := m.spinner.Update(msg)
	cmds := []tea.Cmd{cmd}
	m.spinner = s
	switch msg := msg.(type) {
	case signals.SetStageMsg:
		if (m.status == NotStarted) && (msg.NewStage == m.Stage) {
			targetDir := build_config.GetBuildManager().TargetDir()
			if targetDir == "" {
				log.Fatal("Attempted to build ULLD in an invalid location.")
			}
			m.status = Running
			cmd = signals.SendBeginInitialTemplateCloneMsg(targetDir)
			cmds = append(cmds, cmd)
			// return m, tea.Batch(cmds...)
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keymap.Quit):
			quitMsg := signals.SetQuittingMessage(nil)
			return m, quitMsg
		}
	case responseMsg:
		if msg {
			m.status = Complete
			return m, m.cloneFinishedMsg()
		}
	case signals.BeginInitialTemplateCloneMsg:
		m.status = Running
		go func() {
			defer func() {
				m.status = Complete
				m.Program.Send(signals.SetStage(build_stages.PreConflictResolveBuild)())
			}()
			stage_clone_template_app.Run(msg.TargetDir)
		}()
		return m, tea.Batch(m.spinner.Spinner.Tick, m.waitForActivity(m.sub))
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.spinner.View()
}

func NewCloneTemplateAppUIModel() Model {
	id := "clone-template-app"
	sub := make(chan bool, 1)
	return Model{
		Id:      id,
		Stage:   build_stages.CloneTemplateAppStage,
		status:  NotStarted,
		sub:     sub,
		spinner: general_spinner.NewModel("Cloning ULLD template app..."),
	}
}
