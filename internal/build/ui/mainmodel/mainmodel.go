package build_main_model

import (
	"fmt"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	choose_wait_or_pick_config_loc "github.com/Uh-little-less-dum/cli/internal/build/ui/chooseWaitOrPickConfigLoc"
	clone_template_app "github.com/Uh-little-less-dum/cli/internal/build/ui/cloneTemplateApp"
	confirm_config_dir_loc "github.com/Uh-little-less-dum/cli/internal/build/ui/confirmConfigDirLoc"
	"github.com/Uh-little-less-dum/cli/internal/build/ui/confirmdir"
	"github.com/Uh-little-less-dum/cli/internal/build/ui/filepicker"
	general_confirm "github.com/Uh-little-less-dum/cli/internal/build/ui/generalConfirm"
	general_select_with_desc "github.com/Uh-little-less-dum/cli/internal/build/ui/generalSelectWithDesc"
	pre_conflict_resolve_build_stream "github.com/Uh-little-less-dum/cli/internal/build/ui/preConflictResolveBuild"
	build_stage_utils "github.com/Uh-little-less-dum/cli/internal/buildStageManagement"
	"github.com/Uh-little-less-dum/cli/internal/keymap"
	fs_utils "github.com/Uh-little-less-dum/cli/internal/utils/fs"
	charm_debug "github.com/Uh-little-less-dum/go-utils/pkg/charm/logMessages"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/Uh-little-less-dum/go-utils/pkg/signals"
	"github.com/spf13/viper"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/go-homedir"
)

type mainModel struct {
	help                      help.Model
	confirmDirModel           confirmdir.Model
	targetDirModel            filepicker.Model
	confirmConfigLocEnv       general_confirm.Model
	cloneTemplateAppModel     clone_template_app.Model
	chooseWaitOrPickConfigLoc general_select_with_desc.Model
	pickConfigFile            filepicker.Model
	preConflictResolveStream  pre_conflict_resolve_build_stream.Model
	Program                   *tea.Program
	targetDir                 string
	quitting                  bool
	exitMsg                   string
	manager                   *build_config.BuildManager
}

func (m mainModel) Init() tea.Cmd {
	// FIX: Fix this once the stream model is working.
	return tea.Batch(
		tea.SetWindowTitle("ULLD Build"),
		signals.SetStage(build_stages.PreConflictResolveBuild),
	)
}

func (m *mainModel) ApplyProgramProp(p *tea.Program) {
	m.Program = p
	m.cloneTemplateAppModel.Program = p
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	charm_debug.LogCharmMessages("/Users/bigsexy/Desktop/Go/projects/ulld/cli/messageLog-main.log", "Main", msg)
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
	// TODO: ToPreviousStageMsg is completely untested. Need to implement the keymap in each model to make it work.
	case signals.ToPreviousStageMsg:
		cmd = m.manager.SendToPreviousStageMsg()
		return m, cmd
	case signals.SetStageMsg:
		// WARN: Stages are likely still being modified elsewhere. Fix this to make sure that all modifications to the active stage flow through one function so this can be implemented reliably.
		build_config.SetActiveStage(msg.NewStage)
		if msg.NewStage == m.confirmConfigLocEnv.Stage {
			m.confirmConfigLocEnv.SetDescription(m.manager.ConfigDirPath)
			m.confirmConfigLocEnv, cmd = m.confirmConfigLocEnv.Update(msg)
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)
		}
	case signals.SubCommandCompleteMsg:
		m.preConflictResolveStream, cmd = m.preConflictResolveStream.Update(msg)
		cmds = append(cmds, cmd)
	// This runs when the filepicker model selects a filepath.
	case signals.SetAcceptedTargetDirMsg:
		m.targetDir = msg.TargetDir
		v := viper.GetViper()
		v.Set(string(viper_keys.TargetDirectory), msg.TargetDir)
		build_config.GetBuildManager().TargetDir = msg.TargetDir
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
			cmd = signals.SetStage(build_stages.PickTargetDirStage)
		}
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	case signals.SetQuittingMsg:
		m.quitting = true
		return m, tea.Quit
	case signals.ExitWithMsg:
		m.exitMsg = msg.ExitMsg
		m.quitting = true
		return m, tea.Quit
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}
	switch m.manager.Stage() {
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
	case m.chooseWaitOrPickConfigLoc.Stage:
		m.chooseWaitOrPickConfigLoc, cmd = m.chooseWaitOrPickConfigLoc.Update(msg)
		cmds = append(cmds, cmd)
	case m.preConflictResolveStream.Stage:
		m.preConflictResolveStream, cmd = m.preConflictResolveStream.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	if m.quitting {
		return fmt.Sprintf("\n%s\n\n", m.exitMsg)
	}
	switch m.manager.Stage() {
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
	case m.preConflictResolveStream.Stage:
		return m.preConflictResolveStream.View()
	}
	return s
}

func InitialMainModel(cfg *build_config.BuildManager) *mainModel {
	homeDir, err := homedir.Dir()

	if err != nil {
		log.Fatal(err)
	}

	val := mainModel{
		help:                      help.New(),
		targetDirModel:            filepicker.NewModel(homeDir, fs_utils.DirOnlyDataType, "Where would you like to build ULLD?", build_stages.PickTargetDirStage),
		confirmDirModel:           confirmdir.NewModel("Do you want to build ULLD in your selected directory?", cfg),
		cloneTemplateAppModel:     clone_template_app.NewCloneTemplateAppUIModel(),
		confirmConfigLocEnv:       confirm_config_dir_loc.NewModel(cfg),
		chooseWaitOrPickConfigLoc: choose_wait_or_pick_config_loc.NewModel(),
		pickConfigFile:            filepicker.NewModel(homeDir, fs_utils.FileOnlyDataType, "Select your config file.", build_stages.PickConfigLoc),
		preConflictResolveStream:  pre_conflict_resolve_build_stream.NewModel(cfg),
		targetDir:                 cfg.TargetDir,
		quitting:                  false,
		exitMsg:                   "No worries.",
		manager:                   cfg,
	}

	return &val
}