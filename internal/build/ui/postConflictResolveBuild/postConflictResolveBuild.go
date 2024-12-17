package post_conflict_resolve_build

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	stage_post_conflict_resolve_build "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/post_conflict_resolve_build"
	sub_command_build_stream "github.com/Uh-little-less-dum/cli/internal/build/ui/subCommandBuildStream"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	stream_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/streamIds"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Stage       build_stages.BuildStage
	streamModel sub_command_build_stream.Model
	cfg         *build_config.BuildManager
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

func NewModel(cfg *build_config.BuildManager) Model {
	return Model{
		cfg:         cfg,
		Stage:       build_stages.PreConflictResolveBuild,
		streamModel: sub_command_build_stream.NewModel(stream_ids.PostConflictResolve, build_stages.ResolvePluginConflicts, stage_post_conflict_resolve_build.GetSubStageTree()),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.streamModel, cmd = m.streamModel.Update(msg)
	cmds := []tea.Cmd{cmd}
	switch msg := msg.(type) {
	case signals.SetStageMsg:
		if (m.streamModel.HasNotRun()) && (msg.NewStage == m.Stage) {
			// m.streamModel, cmd = m.streamModel.Run()
			cmds = append(cmds, signals.SendRunSubCommandStreamMsg(m.streamModel.StreamId))
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keymap.Quit):
			quitMsg := signals.SetQuittingMessage(nil)
			return m, quitMsg
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.streamModel.View()
}
