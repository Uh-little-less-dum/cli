package build_main_model

import (
	buildConfig "github.com/Uh-little-less-dum/cli/internal/build/config"
	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	choose_wait_or_pick_config_loc "github.com/Uh-little-less-dum/cli/internal/build/ui/chooseWaitOrPickConfigLoc"
	clone_template_app "github.com/Uh-little-less-dum/cli/internal/build/ui/cloneTemplateApp"
	confirm_config_dir_loc "github.com/Uh-little-less-dum/cli/internal/build/ui/confirmConfigDirLoc"
	"github.com/Uh-little-less-dum/cli/internal/build/ui/confirmdir"
	"github.com/Uh-little-less-dum/cli/internal/build/ui/filepicker"
	general_confirm "github.com/Uh-little-less-dum/cli/internal/build/ui/generalConfirm"
	general_select_with_desc "github.com/Uh-little-less-dum/cli/internal/build/ui/generalSelectWithDesc"
	build_stage_utils "github.com/Uh-little-less-dum/cli/internal/buildStageManagement"
	"github.com/Uh-little-less-dum/cli/internal/keymap"
	"github.com/Uh-little-less-dum/cli/internal/signals"
	fs_utils "github.com/Uh-little-less-dum/cli/internal/utils/fs"
	"github.com/spf13/viper"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/go-homedir"
)

type mainModel struct {
	stage                     constants.BuildStage
	help                      help.Model
	confirmDirModel           confirmdir.Model
	targetDirModel            filepicker.Model
	confirmConfigLocEnv       general_confirm.Model
	cloneTemplateAppModel     clone_template_app.Model
	chooseWaitOrPickConfigLoc general_select_with_desc.Model
	pickConfigFile            filepicker.Model
	targetDir                 string
	quitting                  bool
	manager                   *build_config.BuildManager
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
		// WARN: Stages are likely still being modified elsewhere. Fix this to make sure that all modifications to the active stage flow through one function so this can be implemented reliably.
		m.manager.SetActiveStage(msg.NewStage)
		// m.confirmConfigLocEnv = confirm_config_dir_loc.NewModel()
		if msg.NewStage == m.confirmConfigLocEnv.Stage {
			// m.confirmConfigLocEnv.SetDescription(viper.GetViper().GetString(string(viper_keys.TargetDirectory)))
			m.confirmConfigLocEnv.SetDescription(m.manager.TargetDir)
		}
		m.confirmConfigLocEnv, cmd = m.confirmConfigLocEnv.Update(msg)
		cmds = append(cmds, cmd)
	// This runs when the filepicker model selects a filepath.
	case signals.SetAcceptedTargetDirMsg:
		m.targetDir = msg.TargetDir
		v := viper.GetViper()
		v.Set(string(viper_keys.TargetDirectory), msg.TargetDir)
		_, newStage := build_stage_utils.GetNextBuildStage()
		cmd := signals.SetStage(newStage)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	// This runs when the selected filepath is confirmed or rejected.
	case signals.SetUseSelectedDirMsg:
		if msg.UseSelectedDir {
			_, newStage := build_stage_utils.GetNextBuildStage()
			cmd = signals.SetStage(newStage)
		} else {
			cmd = signals.SetStage(constants.PickTargetDirStage)
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
	case m.cloneTemplateAppModel.Stage:
		m.cloneTemplateAppModel, cmd = m.cloneTemplateAppModel.Update(msg)
		cmds = append(cmds, cmd)
	case m.confirmConfigLocEnv.Stage:
		m.confirmConfigLocEnv, cmd = m.confirmConfigLocEnv.Update(msg)
		cmds = append(cmds, cmd)
	case m.cloneTemplateAppModel.Stage:
		m.cloneTemplateAppModel, cmd = m.cloneTemplateAppModel.Update(msg)
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
	case m.chooseWaitOrPickConfigLoc.Stage:
		return m.chooseWaitOrPickConfigLoc.View()
	case m.pickConfigFile.Stage:
		return m.pickConfigFile.View()
	}
	return s
}

func InitialMainModel(cfg *buildConfig.BuildManager) *mainModel {
	homeDir, err := homedir.Dir()

	if err != nil {
		log.Fatal(err)
	}

	val := mainModel{
		stage:                     cfg.InitialStage(),
		help:                      help.New(),
		targetDirModel:            filepicker.NewModel(homeDir, fs_utils.DirOnlyDataType, "Where would you like to build ULLD?", constants.PickTargetDirStage),
		confirmDirModel:           confirmdir.NewModel("Do you want to build ULLD in your selected directory?", cfg),
		cloneTemplateAppModel:     clone_template_app.NewCloneTemplateAppUIModel(),
		confirmConfigLocEnv:       confirm_config_dir_loc.NewModel(),
		chooseWaitOrPickConfigLoc: choose_wait_or_pick_config_loc.NewModel(),
		pickConfigFile:            filepicker.NewModel(homeDir, fs_utils.FileOnlyDataType, "Select your config file.", constants.PickConfigLoc),
		targetDir:                 cfg.TargetDir,
		quitting:                  false,
		manager:                   cfg,
	}

	return &val
}
