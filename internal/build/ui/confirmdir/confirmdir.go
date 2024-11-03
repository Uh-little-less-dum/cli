package confirmDir

import (
	"fmt"
	"ulld/cli/internal/build/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type (
	errMsg error
)

type model struct {
	stage        constants.BuildStage
	confirmInput *huh.Confirm
	value        *bool
	err          error
}

func InitialModel() model {
	ci := huh.NewConfirm()
	initialVal := false
	modelData := model{
		stage:        constants.ConfirmCurrentDirStage,
		confirmInput: ci,
		value:        &initialVal,
	}

	return modelData
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	ci, cmd := m.confirmInput.Update(msg)
	m.confirmInput = ci.(*huh.Confirm)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.confirmInput.View(),
		"(esc to quit)",
	) + "\n"
}
