package signals

import (
	"github.com/Uh-little-less-dum/cli/internal/build/constants"

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

type SetConfigLocMethod_WaitForClone struct {
}

func SendConfigLocMethod_WaitForClone() tea.Cmd {
	return func() tea.Msg {
		return SetConfigLocMethod_WaitForClone{}
	}
}

type ToPreviousStageMsg struct {
}

func SendToPreviousStageMsg() tea.Cmd {
	return func() tea.Msg {
		return ToPreviousStageMsg{}
	}
}

type SetConfigLocMethod_pickFile struct {
}

func SendConfigLocMethod_pickFile() tea.Cmd {
	return func() tea.Msg {
		return SetConfigLocMethod_pickFile{}
	}
}

func SendUseEnvConfigResponse(wasAccepted bool) tea.Cmd {
	return func() tea.Msg {
		if wasAccepted {
			return SetStageMsg{
				err:      nil,
				NewStage: constants.CloneTemplateAppStage,
			}
		} else {
			return SetStageMsg{
				err:      nil,
				NewStage: constants.ChooseWaitOrPickConfigLoc,
			}
		}
	}
}

type BeginInitialTemplateCloneMsg struct {
	TargetDir string
}

// func openEditor() tea.Cmd {
// 	c := exec.Command(editor) //nolint:gosec
// 	return tea.ExecProcess(c, func(err error) tea.Msg {
// 		return editorFinishedMsg{err}
// 	})
// }

func SendBeginInitialTemplateCloneMsg(targetDir string) tea.Cmd {
	return func() tea.Msg {
		return BeginInitialTemplateCloneMsg{
			TargetDir: targetDir,
		}
	}
}

// func SendBeginInitialTemplateCloneMsg(targetDir string) tea.Cmd {
// 	return func() tea.Msg {
// 		return BeginInitialTemplateCloneMsg{
// 			TargetDir: targetDir,
// 		}
// 	}
// }

type FinishInitialTemplateCloneMsg struct {
}

func SendFinishInitialTemplateCloneMsg() tea.Msg {
	return FinishInitialTemplateCloneMsg{}
}
