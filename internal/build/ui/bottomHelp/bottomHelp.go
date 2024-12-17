package buttom_help

import (
	"strings"

	keyMap "github.com/Uh-little-less-dum/cli/internal/build/keymap"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	keys       keyMap.KeyMap
	help       help.Model
	inputStyle lipgloss.Style
	lastKey    string
	quitting   bool
}

// TODO: Come back here and implement this model, attach a KeyMap object to each model, and create a SetKeys method here to update the keys each time the active model changes.
func NewModel(keys keyMap.KeyMap) Model {
	return Model{
		keys:       keys,
		help:       help.New(),
		inputStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
	}

	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	helpView := m.help.View(m.keys)
	height := 8 - strings.Count(helpView, "\n")

	return "\n" + strings.Repeat("\n", height) + helpView
}
