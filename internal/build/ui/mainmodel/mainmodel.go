package mainBuildModel

import (
	buildConfig "github.com/Uh-little-less-dum/go-utils/pkg/config"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	choose_wait_or_pick_config_loc "github.com/igloo1505/ulldCli/internal/build/ui/chooseWaitOrPickConfigLoc"
	clone_template_app "github.com/igloo1505/ulldCli/internal/build/ui/cloneTemplateApp"
	confirm_config_dir_loc "github.com/igloo1505/ulldCli/internal/build/ui/confirmConfigDirLoc"
	"github.com/igloo1505/ulldCli/internal/build/ui/confirmdir"
	"github.com/igloo1505/ulldCli/internal/build/ui/filepicker"
	general_confirm "github.com/igloo1505/ulldCli/internal/build/ui/generalConfirm"
	general_select_with_desc "github.com/igloo1505/ulldCli/internal/build/ui/generalSelectWithDesc"
	pre_conflict_resolve_build_stream "github.com/igloo1505/ulldCli/internal/build/ui/preConflictResolveBuild"
	stage_gather_config_location "github.com/igloo1505/ulldCli/internal/buildScript/stages/gather_config_location"
	"github.com/igloo1505/ulldCli/internal/keymap"
	fs_utils "github.com/igloo1505/ulldCli/internal/utils/fs"
	"github.com/spf13/viper"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/go-homedir"
)

type mainModel struct {
	stage                         build_stages.BuildStage
	help                          help.Model
	confirmDirModel               confirmdir.Model
	targetDirModel                filepicker.Model
	confirmConfigLocEnv           general_confirm.Model
	cloneTemplateAppModel         clone_template_app.Model
	chooseWaitOrPickConfigLoc     general_select_with_desc.Model
	pickConfigFile                filepicker.Model
	preConflictResolveBuildStream pre_conflict_resolve_build_stream.Model
	targetDir                     string
	quitting                      bool
}

func (m mainModel) Init() tea.Cmd {
	tea.SetWindowTitle("ULLD Build")
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// All list models must receive the size message before they are rendered.
		m.targetDirModel, cmd = m.targetDirModel.Update(msg)
		cmds = append(cmds, cmd)
		m.confirmDirModel, cmd = m.confirmDirModel.Update(msg)
		cmds = append(cmds, cmd)
		m.cloneTemplateAppModel, cmd = m.cloneTemplateAppModel.Update(msg)
		cmds = append(cmds, cmd)
		m.chooseWaitOrPickConfigLoc, cmd = m.chooseWaitOrPickConfigLoc.Update(msg)
		cmds = append(cmds, cmd)
		m.pickConfigFile, cmd = m.pickConfigFile.Update(msg)
		cmds = append(cmds, cmd)
	case signals.SetStageMsg:
		// m.confirmConfigLocEnv = confirm_config_dir_loc.NewModel()
		if msg.NewStage == m.confirmConfigLocEnv.Stage {
			m.confirmConfigLocEnv.SetDescription(viper.GetViper().GetString("appConfigPath"))
			// m.confirmConfigLocEnv = confirm_config_dir_loc.NewModel() // Required to re-read viper based description.
		}
		m.stage = msg.NewStage
		m.confirmConfigLocEnv, cmd = m.confirmConfigLocEnv.Update(msg)
		cmds = append(cmds, cmd)
	case signals.SetAcceptedTargetDirMsg:
		m.targetDir = msg.TargetDir
		v := viper.GetViper()
		v.Set("targetDir", msg.TargetDir)
		_, newStage := stage_gather_config_location.GetNextBuildStage()
		cmd := signals.SetStage(newStage)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	case signals.SetUseSelectedDirMsg:
		if msg.UseSelectedDir {
			_, newStage := stage_gather_config_location.GetNextBuildStage()
			cmd = signals.SetStage(newStage)
		} else {
			cmd = signals.SetStage(build_stages.PickTargetDirStage)
		}
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	case signals.SetQuittingMsg:
		m.quitting = true
		return m, tea.Quit
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}
	switch m.stage {
	case m.confirmDirModel.Stage:
		m.confirmDirModel, cmd = m.confirmDirModel.Update(msg)
		cmds = append(cmds, cmd)
	case m.targetDirModel.Stage:
		m.targetDirModel, cmd = m.targetDirModel.Update(msg)
		cmds = append(cmds, cmd)
	case m.confirmConfigLocEnv.Stage:
		m.confirmConfigLocEnv, cmd = m.confirmConfigLocEnv.Update(msg)
		cmds = append(cmds, cmd)
	case m.cloneTemplateAppModel.Stage:
		m.cloneTemplateAppModel, cmd = m.cloneTemplateAppModel.Update(msg)
		cmds = append(cmds, cmd)
	case m.preConflictResolveBuildStream.Stage:
		m.preConflictResolveBuildStream, cmd = m.preConflictResolveBuildStream.Update(msg)
		cmds = append(cmds, cmd)
	case m.chooseWaitOrPickConfigLoc.Stage:
		m.chooseWaitOrPickConfigLoc, cmd = m.chooseWaitOrPickConfigLoc.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	if m.quitting {
		return "\n  No worries.\n\n"
	}
	switch m.stage {
	case m.confirmDirModel.Stage:
		return m.confirmDirModel.View()
	case m.targetDirModel.Stage:
		return m.targetDirModel.View()
	case m.confirmConfigLocEnv.Stage:
		return m.confirmConfigLocEnv.View()
	case m.cloneTemplateAppModel.Stage:
		return m.cloneTemplateAppModel.View()
	case m.preConflictResolveBuildStream.Stage:
		return m.preConflictResolveBuildStream.View()
	case m.chooseWaitOrPickConfigLoc.Stage:
		return m.chooseWaitOrPickConfigLoc.View()
	case m.pickConfigFile.Stage:
		return m.pickConfigFile.View()
	}
	return s
}

func InitialMainModel(cfg *buildConfig.BuildConfigOpts) *mainModel {
	homeDir, err := homedir.Dir()

	if err != nil {
		log.Fatal(err)
	}

	val := mainModel{
		stage:                         cfg.InitialStage,
		help:                          help.New(),
		targetDirModel:                filepicker.NewModel(homeDir, fs_utils.DirOnlyDataType, "Where would you like to build ULLD?", build_stages.PickTargetDirStage),
		confirmDirModel:               confirmdir.NewModel("Do you want to build ULLD in your selected directory?"),
		cloneTemplateAppModel:         clone_template_app.NewCloneTemplateAppUIModel(),
		preConflictResolveBuildStream: pre_conflict_resolve_build_stream.NewModel(cfg),
		confirmConfigLocEnv:           confirm_config_dir_loc.NewModel(),
		chooseWaitOrPickConfigLoc:     choose_wait_or_pick_config_loc.NewModel(),
		pickConfigFile:                filepicker.NewModel(homeDir, fs_utils.FileOnlyDataType, "Select your config file.", build_stages.PickConfigLoc),
		targetDir:                     cfg.TargetDir,
		quitting:                      false,
	}

	return &val
}
