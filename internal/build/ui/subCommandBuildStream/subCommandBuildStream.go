package sub_command_build_stream

import (
	"fmt"
	"strings"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	run_status "github.com/Uh-little-less-dum/go-utils/pkg/constants/runStatus"
	stream_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/streamIds"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	subStages []*sub_stage.SubStage
	index     int
	width     int
	height    int
	spinner   spinner.Model
	progress  progress.Model
	StreamId  stream_ids.StreamId
	Status    run_status.RunStatus
	nextStage build_stages.BuildStage
	cfg       *build_config.BuildManager
}

var (
	doneStyle = lipgloss.NewStyle().Margin(1, 2)
)

func NewModel(streamId stream_ids.StreamId, nextStage build_stages.BuildStage, subCommands []*sub_stage.SubStage) Model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return Model{
		subStages: subCommands,
		spinner:   s,
		progress:  p,
		Status:    run_status.NotStarted,
		index:     0,
		nextStage: nextStage,
		cfg:       build_config.GetBuildManager(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case sub_stage.InstalledPkgMsg:
		pkg := m.subStages[m.index]
		cm := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
		if m.index >= len(m.subStages)-1 {
			// Everything's been installed. We're done!
			m.Status = run_status.Complete
			return m, tea.Sequence(
				tea.Printf("%s %s", cm, pkg.Name), // print the last success message
				tea.Quit,                          // exit the program
			)
		}

		// Update progress bar
		m.index++
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.subStages)))

		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", cm, pkg.Name), // print success message above our program
			sub_stage.DownloadAndInstall(m.subStages[m.index]), // download the next package
		)
	case signals.RunSubCommandStreamMsg:
		if msg.StreamId == m.StreamId {
			return m.Run()
		}
	case signals.SubCommandCompleteMsg:
		c := m.subStages[m.index]
		if m.index >= len(m.subStages)-1 {
			// Everything's been installed. We're done!
			m.Status = run_status.Complete
			return m, tea.Sequence(
				tea.Println(c.CompleteUserMessage()), // print the last success message
				m.OnComplete(),
			)
		}

		// Update progress bar
		m.index++
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.subStages)))

		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s", c.CompleteUserMessage()), // print success message above our program
			m.runSubCommand(m.subStages[m.index]),     // download the next package
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m Model) OnComplete() tea.Cmd {
	return signals.SetStage(m.nextStage)
}

func (m *Model) Run() (Model, tea.Cmd) {
	return *m, tea.Batch(m.runSubCommand(m.subStages[m.index]), m.spinner.Tick)
}

func (m Model) View() string {
	n := len(m.subStages)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.Status == run_status.Complete {
		return doneStyle.Render("Done!\n")
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	// pkgName := currentPkgNameStyle.Render()
	// info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Installing " + pkgName)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render(m.subStages[m.index].InProgressUserMessage())

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

func (m *Model) runSubCommand(subCommand *sub_stage.SubStage) tea.Cmd {
	return subCommand.Run(m.cfg, m.StreamId)
}

func (m Model) HasNotRun() bool {
	return run_status.HasNotRun(m.Status)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
