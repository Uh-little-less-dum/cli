package resolve_plugin_conflicts

import (
	"fmt"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	cli_styles "github.com/Uh-little-less-dum/go-utils/pkg/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form      *huh.Form
	Stage     build_stages.BuildStage
	nextStage build_stages.BuildStage
	confirm   *huh.Confirm
	cfg       *build_config.BuildManager
}

const formKey string = "useCwd"

func NewModel(buildManager *build_config.BuildManager) Model {
	c := huh.NewConfirm().
		Key(formKey).
		Title(fmt.Sprintf("We found %d conflicts.", len(buildManager.PluginConflicts()))).
		Affirmative("Yup").
		Negative("No")
	c.Description("Don't worry. Overlap is expected when using external plugins.")
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				c,
			),
		).WithTheme(cli_styles.GetHuhTheme()),
		Stage:     build_stages.ConfirmCurrentDirStage,
		confirm:   c,
		cfg:       buildManager,
		nextStage: build_stages.PostConflictResolveBuild,
	}
}

func (m Model) Init() tea.Cmd {

	if len(m.cfg.PluginConflicts()) == 0 {
		return signals.SetStage(m.nextStage)
	}

	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

	// if build_config.IsActiveStage(m.Stage) && build_config.ShouldSkipStage(m.Stage) {
	// 	return m, signals.SetUseSelectedDir(true)
	// }

	form, cmd := m.form.Update(msg)

	var cmds = []tea.Cmd{cmd}

	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	switch msgType := msg.(type) {
	case signals.SetStageMsg:
		if msgType.NewStage == m.Stage {
			cmds = append(cmds, m.confirm.Focus())
		} else {
			cmds = append(cmds, m.confirm.Blur())
		}
	}

	if m.form.State == huh.StateCompleted {
		d := m.form.GetBool(formKey)
		c := signals.SetUseSelectedDir(d)
		cmds = append(cmds, c)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	// TODO: Come back and add a description or subtitle showing the currently populated active directory if we have that field and it's not empty.
	return "\n" + m.form.View()
}
