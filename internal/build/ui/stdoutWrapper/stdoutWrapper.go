package stdout_wrapper

import (
	"bufio"
	"io"
	"os/exec"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/igloo1505/ulldCli/internal/signals"
)

const viewportHeight = 8

type outputMsg string

type outputErrMsg struct{ err error }

type outputDoneMsg struct{}

func ping(domain string, sub chan string) tea.Cmd {
	return func() tea.Msg {
		// Setup the command
		cmd := exec.Command("ping", "-c", "10", domain)
		out, err := cmd.StdoutPipe()
		if err != nil {
			return outputErrMsg{err}
		}

		// Run the command.
		if err := cmd.Start(); err != nil {
			return outputErrMsg{err}
		}

		// Read command output as it arrives.
		buf := bufio.NewReader(out)
		for {
			line, _, err := buf.ReadLine()
			if err == io.EOF {
				return outputDoneMsg{}
			}
			if err != nil {
				return outputErrMsg{err}
			}
			// Send output to our program.
			sub <- string(line)
		}
	}
}

func waitForPingResponses(sub chan string) tea.Cmd {
	return func() tea.Msg {
		// Send the ping to Update.
		return outputMsg(<-sub)
	}
}

type Model struct {
	sub      chan string
	result   *string
	viewport viewport.Model
}

func NewModel(initialString string) Model {
	m := Model{
		result:   new(string),
		sub:      make(chan string),
		viewport: viewport.New(0, viewportHeight),
	}
	m.appendOutput(initialString)
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(ping("charm.sh", m.sub), waitForPingResponses(m.sub))
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case signals.StdOutWrapperOutputMsg:
		m.appendOutput(string(msg.Body))
		cmds = append(cmds, waitForPingResponses(m.sub))
	case outputErrMsg:
		m.appendOutput("Error: " + msg.err.Error())
	case outputDoneMsg:
		m.appendOutput("Done. Press ^C to exit.")
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) appendOutput(s string) {
	*m.result += "\n" + s
	m.viewport.SetContent(*m.result)
	m.viewport.GotoBottom()
}

func (m Model) View() string {
	return m.viewport.View()
}

func (m Model) Write(p []byte) (int, error) {
	signals.SendStdOutWrapperOutputMsg(string(p))
	return len(p), nil
}
