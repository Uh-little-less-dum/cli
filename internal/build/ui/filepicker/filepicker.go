package filepicker

import (
	"ulld/cli/internal/build/constants"
	cli_styles "ulld/cli/internal/styles"
	fs_utils "ulld/cli/internal/utils/fs"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
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
	enterItem        key.Binding
	goToParent       key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		enterItem: key.NewBinding(
			key.WithKeys("o"),
			key.WithHelp("o", "enter directory"),
		),
		goToParent: key.NewBinding(
			key.WithKeys("O"),
			key.WithHelp("O", "to parent directory"),
		),
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
	fsDir         *fs_utils.FSDirectory
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
	case SetNewDirMessage:
		items := getParentListItems(m.fsDir)
		setItemsCmd := m.list.SetItems(items)
		return m, setItemsCmd
	case SetParentDirMessage:
		items := getListItems(m.fsDir, msg.NewDir, false)
		setItemsCmd := m.list.SetItems(items)
		return m, setItemsCmd
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

		case key.Matches(msg, m.keys.enterItem):
			newItem := m.list.SelectedItem().FilterValue()
			cmd := SetNewFilePickerDir(newItem)
			statusCmd := m.list.NewStatusMessage(statusMessageStyle("Now in " + newItem))
			return m, tea.Batch(cmd, statusCmd)
		case key.Matches(msg, m.keys.goToParent):
			cmd := SetParentDir()
			return m, cmd
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

func getParentListItems(fsDir *fs_utils.FSDirectory) []list.Item {
	itemStrings := fsDir.GetParentData()
	numItems := len(itemStrings)
	items := make([]list.Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = item{title: itemStrings[i], desc: itemStrings[i]}
	}
	return items
}

func getListItems(fsDir *fs_utils.FSDirectory, dirPath string, isInitial bool) []list.Item {
	var itemStrings []string
	if isInitial {
		itemStrings = fsDir.GetDataFromAbsolutePath(dirPath)
	} else {
		itemStrings = fsDir.GetNewData(dirPath)
	}
	numItems := len(itemStrings)
	items := make([]list.Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = item{title: itemStrings[i], desc: itemStrings[i]}
	}
	return items
}

func NewModel(initialDir string, dataType fs_utils.FilePickerDataType, title string) Model {
	fsDir := fs_utils.FSDirectory{Path: initialDir, DataType: dataType}
	items := getListItems(&fsDir, initialDir, true)
	delegateKeys := newDelegateKeyMap()
	delegate := newItemDelegate(delegateKeys)
	delegate.Styles.SelectedTitle = selectedTitleStyle
	delegate.Styles.SelectedDesc = selectedTitleStyle.Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"})

	listKeys := newListKeyMap()

	list := list.New(items, delegate, 0, 0)
	list.Styles.Title = cli_styles.TitleStyle
	list.Title = title
	list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.enterItem,
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
		Stage:        constants.PickTargetDirStage,
		state:        initializing,
		fsDir:        &fsDir,
	}
}
