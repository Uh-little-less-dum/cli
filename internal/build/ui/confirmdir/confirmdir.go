package confirmdir

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	cli_styles "github.com/Uh-little-less-dum/go-utils/pkg/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form    *huh.Form
	Stage   build_stages.BuildStage
	confirm *huh.Confirm
}

const formKey string = "useCwd"

func NewModel(title string, buildManager *build_config.BuildManager) Model {
	theme := cli_styles.GetHuhTheme()
	c := huh.NewConfirm().
		Key(formKey).
		Title(title).
		Affirmative("Yup").
		Negative("No")

	if buildManager.TargetDir != "" {
		c.Description(buildManager.TargetDir)
	}
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				c,
			),
		).WithTheme(theme),
		Stage:   build_stages.ConfirmCurrentDirStage,
		confirm: c,
	}
}

func (m Model) Init() tea.Cmd {
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
		if msgType.NewStage == build_stages.ConfirmCurrentDirStage {
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
