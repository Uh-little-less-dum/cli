package confirmDir

import (
	"fmt"
	"ulld/cli/internal/build/constants"
	keyMap "ulld/cli/internal/build/keymap"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type (
	errMsg error
)

type ConfirmCurrentDirModel struct {
	Stage        constants.BuildStage
	confirmInput *huh.Confirm
	keys         keyMap.KeyMap
	value        *bool
	err          error
}

func InitialModel(title string, desc string) *ConfirmCurrentDirModel {
	ci := huh.NewConfirm()
	ci.Title(title)
	ci.Description(desc)
	initialVal := false
	m := ConfirmCurrentDirModel{
		Stage:        constants.ConfirmCurrentDirStage,
		confirmInput: ci,
		value:        &initialVal,
		keys:         keyMap.DefaultKeymap,
	}

	return &m
}

func (m ConfirmCurrentDirModel) Init() tea.Cmd {
	return nil
}

func (m *ConfirmCurrentDirModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
			// case key.Matches(msg, m.keys.Up):
			//
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	ci, cmd := m.confirmInput.Update(msg)

	// WARN: This ci.(*huh.Confirm) is printing 'no'. Not sure if it just prints that way, or if it's actually returning 'no'. Look up docs on type castng in Go ASAP.
	m.confirmInput = ci.(*huh.Confirm)

	return m, cmd
}

func (m *ConfirmCurrentDirModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.confirmInput.View(),
		"(esc to quit)",
	) + "\n"
}
