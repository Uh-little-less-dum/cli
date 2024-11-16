package clone_template_app

import (
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/igloo1505/ulldCli/internal/build/constants"
	clone_template_app_manager "github.com/igloo1505/ulldCli/internal/build/stages"
	"github.com/igloo1505/ulldCli/internal/build/ui/progressbar"
	"github.com/spf13/viper"
)

type Model struct {
	progress progressbar.Model
	Id       string
	Stage    constants.BuildStage
	status   int
	err      error
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

//	func (m Model) View() string {
//		return m.progress.View()
//	}
func (m Model) View() string {
	targetDir := viper.GetViper().GetString("targetDir")
	s := fmt.Sprintf("Cloning %s into %s", constants.SparseCloneRepoUrl, targetDir)
	if m.err != nil {
		s += fmt.Sprintf("something went wrong: %s", m.err)
	} else if m.status != 0 {
		s += fmt.Sprintf("%d %s", m.status, http.StatusText(m.status))
	}
	return s + "\n"
}

func (m Model) Write(p []byte) (n int, err error) {
	return 0, nil
}

func beginSparseClone(targetDir string) {
	clone_template_app_manager.CloneTemplateApp(targetDir)
}

func NewCloneTemplateAppUIModel() Model {
	id := "clone-template-app"
	return Model{
		progress: progressbar.NewModel(id),
		Id:       id,
		Stage:    constants.CloneTemplateAppStage,
	}
}
