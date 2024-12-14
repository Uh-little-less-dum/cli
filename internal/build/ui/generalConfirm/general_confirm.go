package general_confirm

import (
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	cli_styles "github.com/Uh-little-less-dum/cli/internal/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const valueKey string = "adbajkafajl"

type OnResponseFunc func(wasAccepted bool) tea.Cmd

type Model struct {
	form        *huh.Form
	Stage       build_stages.BuildStage
	confirm     *huh.Confirm
	Description string
	OnResponse  OnResponseFunc
	Value       *bool
}

func NewModel(title, desc string, onResponse OnResponseFunc, stage build_stages.BuildStage) Model {
	theme := cli_styles.GetHuhTheme()
	var b bool
	c := huh.NewConfirm().
		Value(&b).
		Key(valueKey).
		Title(title).
		Affirmative("Yes").
		Negative("No")

	if desc != "" {
		c = c.Description(desc)
	}
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				c,
			),
		).WithTheme(theme),
		OnResponse: onResponse,
		Stage:      stage,
		confirm:    c,
		Value:      &b,
	}
}

func (m Model) Focus() tea.Cmd {
	return m.confirm.Focus()
}

type forceUpdate struct{}

func sendForceUpdate() tea.Msg {
	return forceUpdate{}
}

func (m *Model) SetDescription(desc string) {
	m.confirm.Description(desc)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.form.Init(), sendForceUpdate)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

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
		d := m.form.GetBool(valueKey)
		cmds = append(cmds, m.OnResponse(d))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	// TODO: Come back and add a description or subtitle showing the currently populated active directory if we have that field and it's not empty.
	return "\n" + m.form.View()
}
