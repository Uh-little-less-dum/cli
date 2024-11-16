package progressbar

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type ProgressPercent struct {
	percent    float32
	progressId string
}

type Model struct {
	progress progress.Model
	Id       string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case ProgressPercent:
		if msg.progressId == m.Id {
			if m.progress.Percent() == 1.0 {
				return m, tea.Quit
			}
			cmd := m.progress.SetPercent(float64(msg.percent))
			return m, cmd
		}

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
	return m, nil
}

func (m Model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func (m Model) SendProgressCmd(progressPercent float32) tea.Cmd {
	return func() tea.Msg {
		return ProgressPercent{
			percent:    progressPercent,
			progressId: m.Id,
		}
	}
}

func NewModel(id string) Model {
	return Model{
		progress: progress.New(progress.WithDefaultGradient()),
		Id:       id,
	}

}
