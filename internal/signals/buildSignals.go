package signals

import (
	"ulld/cli/internal/build/constants"

	tea "github.com/charmbracelet/bubbletea"
)

// type errMsg struct{ error }

type SetStageMsg struct {
	err      error
	NewStage constants.BuildStage
}

func SetStage(newStage constants.BuildStage) tea.Cmd {
	return func() tea.Msg {
		return SetStageMsg{
			err:      nil,
			NewStage: newStage,
		}
	}
}

type SetUseSelectedDirMsg struct {
	err            error
	UseSelectedDir bool
}

func SetUseSelectedDir(shouldUse bool) tea.Cmd {
	return func() tea.Msg {
		return SetUseSelectedDirMsg{
			err:            nil,
			UseSelectedDir: shouldUse,
		}
	}
}
