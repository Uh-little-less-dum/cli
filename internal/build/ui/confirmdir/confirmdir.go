package confirmdir

import (
	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	"github.com/Uh-little-less-dum/cli/internal/signals"
	cli_styles "github.com/Uh-little-less-dum/go-utils/pkg/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form    *huh.Form
	Stage   constants.BuildStage
	confirm *huh.Confirm
}

const formKey string = "useCwd"

func NewModel(title string, buildManager *build_config.BuildManager) Model {
	theme := cli_styles.GetHuhTheme()
	c := huh.NewConfirm().
		Key(formKey).
		Title(title).
		Affirmative("Yup").
		Negative("No")

	if buildManager.TargetDir != "" {

		c.Description(buildManager.TargetDir)
	}

	// d := viper.GetViper().GetString(string(viper_keys.TargetDirectory))
	// if d != "" {
	// 	c.Description(d)
	// } else {
	// 	cwd, err := os.Getwd()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	c.Description(cwd)
	// }
	return Model{
		form: huh.NewForm(
			huh.NewGroup(
				c,
			),
		).WithTheme(theme),
		Stage:   constants.ConfirmCurrentDirStage,
		confirm: c,
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

	// if build_config.IsActiveStage(m.Stage) && build_config.ShouldSkipStage(m.Stage) {
	// 	return m, signals.SetUseSelectedDir(true)
	// }

	form, cmd := m.form.Update(msg)

	var cmds = []tea.Cmd{cmd}

	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	switch msgType := msg.(type) {
	case signals.SetStageMsg:
		if msgType.NewStage == constants.ConfirmCurrentDirStage {
			cmds = append(cmds, m.confirm.Focus())
		} else {
			cmds = append(cmds, m.confirm.Blur())
		}
	}

	if m.form.State == huh.StateCompleted {
		d := m.form.GetBool(formKey)
		c := signals.SetUseSelectedDir(d)
		cmds = append(cmds, c)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	// TODO: Come back and add a description or subtitle showing the currently populated active directory if we have that field and it's not empty.
	return "\n" + m.form.View()
}
