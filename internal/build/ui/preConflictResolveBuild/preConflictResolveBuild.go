package pre_conflict_resolve_build_stream

import (
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	build_config "github.com/Uh-little-less-dum/go-utils/pkg/config"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	run_status "github.com/Uh-little-less-dum/go-utils/pkg/constants/runStatus"
	stream_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/streamIds"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	sub_command_build_stream "github.com/igloo1505/ulldCli/internal/build/ui/subCommandBuildStream"
	"github.com/spf13/viper"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type Model struct {
	Stage       build_stages.BuildStage
	streamModel sub_command_build_stream.Model
	status      run_status.RunStatus
	cfg         *build_config.BuildConfigOpts
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

func getSubStages() []*sub_stage.SubStage {
	data := []*sub_stage.SubStage{}
	return data
}

func NewModel(cfg *build_config.BuildConfigOpts) Model {
	return Model{
		cfg:         cfg,
		Stage:       build_stages.PreConflictResolveBuild,
		streamModel: sub_command_build_stream.NewModel(stream_ids.PreConflictResolve, build_stages.ConfirmWaitForConfigMove, getSubStages()),
	}
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
			m.status = run_status.Running
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keymap.Quit):
			quitMsg := signals.SetQuittingMessage(nil)
			return m, quitMsg
		}
	}
	return m, nil
}

func (m Model) beginPreConflictResolveBuild(targetDir string) {
	log.Fatal("Removed this stage wrapper. Implement the build_sub_stage stream model here.")
	// stage_pre_conflict_resolve_build.PreConflictResolveBuild(targetDir)
}

func (m Model) View() string {
	return m.streamModel.View()
}
