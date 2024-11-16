package clone_template_app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/igloo1505/ulldCli/internal/build/constants"
	clone_template_app_manager "github.com/igloo1505/ulldCli/internal/build/stages"
	"github.com/igloo1505/ulldCli/internal/build/ui/progressbar"
	"github.com/igloo1505/ulldCli/internal/signals"
	"github.com/spf13/viper"
)

type CloneStatus int

const (
	NotStarted CloneStatus = iota
	Running
	Complete
)

type Model struct {
	progress progressbar.Model
	Id       string
	Stage    constants.BuildStage
	status   CloneStatus
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	log.Fatal("msg", msg)
	switch msg.(type) {
	case signals.SetStageMsg:
		if m.status == NotStarted {
			targetDir := viper.GetViper().GetString("targetDir")
			if targetDir == "" {
				log.Fatal("Attempted to build ULLD in an invalid location.")
			}
			m.beginSparseClone(targetDir)
			m.status = Running
		}
	}
	return m, nil
}

//	func (m Model) View() string {
//		return m.progress.View()
//	}

func (m Model) View() string {
	targetDir := viper.GetViper().GetString("targetDir")
	s := fmt.Sprintf("Cloning %s into %s", constants.SparseCloneRepoUrl, targetDir)
	// if m.err != nil {
	// 	s += fmt.Sprintf("something went wrong: %s", m.err)
	// }
	return s + "\n"
}

func (m Model) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (m Model) beginSparseClone(targetDir string) {
	clone_template_app_manager.CloneTemplateApp(targetDir)
}

func NewCloneTemplateAppUIModel() Model {
	id := "clone-template-app"
	return Model{
		progress: progressbar.NewModel(id),
		Id:       id,
		Stage:    constants.CloneTemplateAppStage,
		status:   NotStarted,
	}
}
