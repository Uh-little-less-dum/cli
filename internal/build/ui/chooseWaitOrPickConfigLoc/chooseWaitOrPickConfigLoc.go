package choose_wait_or_pick_config_loc

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	general_select_with_desc "github.com/Uh-little-less-dum/cli/internal/build/ui/generalSelectWithDesc"
	"github.com/Uh-little-less-dum/cli/internal/signals"
)

const (
	selectFileTitle   string = "Select a file"
	waitForCloneTitle string = "Manually add your config"
)

func sendConfigLocMethod_handlePickOrWait(acceptedTitle string) tea.Cmd {
	return func() tea.Msg {
		if acceptedTitle == selectFileTitle {
			return signals.SetStageMsg{
				NewStage: constants.PickConfigLoc,
			}
		} else {
			return signals.SetStageMsg{
				NewStage: constants.PickConfigLoc,
			}
		}
	}
}

func NewModel() general_select_with_desc.Model {
	selectFile := general_select_with_desc.Item{}
	selectFile.SetTitle(selectFileTitle)
	selectFile.SetDescription("This will open a file picker so you can find your configuration file.")

	waitForClone := general_select_with_desc.Item{}
	waitForClone.SetTitle(waitForCloneTitle)
	waitForClone.SetDescription("We can clone the initial template and then pause for you to move your file.")

	opts := []general_select_with_desc.Item{
		selectFile,
		waitForClone,
	}

	m := general_select_with_desc.NewModel(opts, "How would you like to proceed?", sendConfigLocMethod_handlePickOrWait, constants.ChooseWaitOrPickConfigLoc)
	return m
}
