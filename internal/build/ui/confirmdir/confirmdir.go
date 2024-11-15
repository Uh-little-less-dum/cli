package confirmdir

import (
	"ulld/cli/internal/build/constants"
	"ulld/cli/internal/signals"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form  *huh.Form
	Stage constants.BuildStage
}

func NewModel(title string) Model {
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Key("useCurrentDir").
					Title(title),
			),
		),
		Stage: constants.ConfirmCurrentDirStage,
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

	if m.form.State == huh.StateCompleted {
		d := m.form.GetBool("useCurrentDir")
		c := signals.SetUseSelectedDir(d)
		cmds = append(cmds, c)
	}

	return m, tea.Batch(cmds...)

}

func (m Model) View() string {
	// TODO: Come back and add a description or subtitle showing the currently populated active directory if it exist.
	return m.form.View()
}
