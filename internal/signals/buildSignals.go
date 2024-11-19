package signals

import (
	"github.com/igloo1505/ulldCli/internal/build/constants"

	tea "github.com/charmbracelet/bubbletea"
)

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

type SetAcceptedTargetDirMsg struct {
	err       error
	TargetDir string
}

func SetAcceptedTargetDir(dirPath string) tea.Cmd {
	return func() tea.Msg {
		return SetAcceptedTargetDirMsg{
			err:       nil,
			TargetDir: dirPath,
		}
	}
}

type SetQuittingMsg struct {
	err error
}

func SetQuittingMessage(err error) tea.Cmd {
	return func() tea.Msg {
		return SetQuittingMsg{
			err: err,
		}
	}
}

type StdOutWrapperOutputMsg struct {
	Body string
}

func SendStdOutWrapperOutputMsg(content string) tea.Cmd {
	return func() tea.Msg {
		return StdOutWrapperOutputMsg{
			Body: content,
		}
	}
}
