package general_select

import (
	"github.com/igloo1505/ulldCli/internal/build/constants"
	cli_styles "github.com/igloo1505/ulldCli/internal/styles"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle           = lipgloss.NewStyle().Padding(1, 2)
	selectedTitleStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(cli_styles.UlldBlueLipgloss).
				Foreground(cli_styles.UlldBlueLipgloss).
				Padding(0, 0, 0, 1)
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type state int

const (
	initializing state = iota
	ready
)

type Model struct {
	list          list.Model
	Stage         constants.BuildStage
	keys          *listKeyMap
	delegateKeys  *delegateKeyMap
	state         state
	width, height int
}

func (m Model) Init() tea.Cmd {
	return nil
}

type SetNewDirMessage struct {
	err    error
	NewDir string
}

func SetNewFilePickerDir(newDir string) tea.Cmd {
	return func() tea.Msg {
		return SetNewDirMessage{
			err:    nil,
			NewDir: newDir,
		}
	}
}

type SetParentDirMessage struct {
	err    error
	NewDir string
}

func SetParentDir() tea.Cmd {
	return func() tea.Msg {
		return SetParentDirMessage{
			err: nil,
		}
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.width, m.height = msg.Width, msg.Height
		m.state = ready
		return m, nil
	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if (m.list.FilterState() == list.Filtering) || (m.state == initializing) {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd

		case key.Matches(msg, m.keys.toggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil

		case key.Matches(msg, m.keys.toggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil

		case key.Matches(msg, m.keys.togglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil

		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil
		}
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	res := m.list.View()
	if m.state == initializing {
		return "initializing..."
	}
	return appStyle.Render(res)
}

func getListItems(opts []string) []list.Item {
	var items []list.Item
	for _, s := range opts {
		items = append(items, item{title: s, desc: s})
	}
	return items
}

func NewModel(opts []string, title string, onAccept OnAcceptFunc, stage constants.BuildStage) Model {
	items := getListItems(opts)
	delegateKeys := newDelegateKeyMap()
	delegate := newItemDelegate(delegateKeys, onAccept)
	delegate.Styles.SelectedTitle = selectedTitleStyle
	delegate.Styles.SelectedDesc = selectedTitleStyle.Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"})

	listKeys := newListKeyMap()

	list := list.New(items, delegate, 0, 0)
	list.Styles.Title = cli_styles.TitleStyle
	list.Title = title
	list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}
	return Model{
		list:         list,
		keys:         listKeys,
		delegateKeys: delegateKeys,
		Stage:        stage,
		state:        initializing,
	}
}
