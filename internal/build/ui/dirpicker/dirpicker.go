package dirPicker

import (
	"errors"
	"os"
	"strings"
	"time"
	"ulld/cli/internal/build/constants"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type DirPickerModel struct {
	filepicker   filepicker.Model
	Stage        constants.BuildStage
	selectedFile string
	quitting     bool
	err          error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m *DirPickerModel) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m DirPickerModel) GetValue() string {
	return m.selectedFile
}

func (m *DirPickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	didSelect, val := m.filepicker.DidSelectFile(msg)
	if didSelect {
		log.Debug(val)
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m DirPickerModel) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Pick a directory:")
	} else {
		s.WriteString("Selected directory: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return s.String()
}

func InitialDirPicker() *DirPickerModel {
	fp := filepicker.New()
	fp.DirAllowed = true
	fp.FileAllowed = false
	fp.AllowedTypes = []string{}
	fp.AutoHeight = true
	fp.CurrentDirectory, _ = os.UserHomeDir()
	m := DirPickerModel{
		filepicker: fp,
		err:        nil,
		Stage:      constants.PickTargetDirStage,
	}
	return &m
}
