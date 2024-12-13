package confirmdir

import (
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	cli_styles "github.com/igloo1505/ulldCli/internal/styles"
	"github.com/spf13/viper"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form    *huh.Form
	Stage   build_stages.BuildStage
	confirm *huh.Confirm
}

func NewModel(title string) Model {
	theme := cli_styles.GetHuhTheme()
	c := huh.NewConfirm().
		Key("useCurrentDir").
		Title(title).
		Affirmative("Yup").
		Negative("No")

	d := viper.GetViper().GetString("targetDir")
	if d != "" {
		c.Description(d)
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
		d := m.form.GetBool("useCurrentDir")
		c := signals.SetUseSelectedDir(d)
		cmds = append(cmds, c)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	// TODO: Come back and add a description or subtitle showing the currently populated active directory if we have that field and it's not empty.
	return "\n" + m.form.View()
}
