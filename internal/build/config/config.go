package build_config

import (
	"os"

	env_vars "github.com/Uh-little-less-dum/build/pkg/envVars"
	"github.com/Uh-little-less-dum/build/pkg/utils"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	form_data "github.com/Uh-little-less-dum/cli/internal/build/formData"
	"github.com/Uh-little-less-dum/cli/internal/signals"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type BuildManager struct {
	form_data.BuildFormData
	AppConfigPath string
	ConfigDirPath string
	LogLevel      string
	LogFile       string
	stack         []constants.BuildStage
	skipStages    []constants.BuildStage
	allowGoBack   bool
}

var b *BuildManager

// Returns the BuildManager singleton.
func GetBuildManager() *BuildManager {
	return b
}

func init() {
	val := BuildManager{
		AppConfigPath: utils.GetDefaultAppConfigPath(),
	}

	if val.TargetDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		val.TargetDir = cwd
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	val.TargetDir = cwd
	// Handle initial values with environment variables here.
	val.ConfigDirPath = os.Getenv(string(env_vars.AdditionalSources))
	val.LogLevel = os.Getenv(string(env_vars.LogLevel))
	val.LogFile = os.Getenv(string(env_vars.LogFile))
	val.stack = []constants.BuildStage{constants.ConfirmCurrentDirStage}
	b = &val
}

func (b *BuildManager) SetInitialStage(s constants.BuildStage) {
	b.stack = []constants.BuildStage{s}
}

// Adds a build stage that should be skipped to the list of skipped stages. This list can then be accessed from within each build stage at which point the stage can be bypassed.
func (b *BuildManager) AddSkippedStage(stage constants.BuildStage) {
	b.skipStages = append(b.skipStages, stage)
}

func (b *BuildManager) RemoveStageFromSkipped(stage constants.BuildStage) {
	var skipStages []constants.BuildStage
	for _, s := range b.skipStages {
		if s != stage {
			skipStages = append(skipStages, s)
		}
	}
	b.skipStages = skipStages
}

// TODO: Set up the rest of the config flags here. Handle what you can with viper first though. Might not even need this struct if viper can handle everything.
func (c *BuildManager) Init(args []string) {
	v := viper.GetViper()
	acPath := v.GetString(string(viper_keys.AppConfigPath))
	hasTargetDir := (len(args) > 0) && (args[0] != "")
	var initialStage constants.BuildStage
	if (v.GetBool(string(viper_keys.UseCwd))) || (hasTargetDir) {
		initialStage = utils.Ternary(acPath != "", constants.CloneTemplateAppStage, constants.ConfirmConfigLocFromEnv)
	} else if acPath != "" {
		initialStage = constants.CloneTemplateAppStage
	} else {
		initialStage = constants.ConfirmCurrentDirStage
	}
	c.SetInitialStage(initialStage)
}

// Utility to check if stage should be skipped. Reads data formerly set by AddSkippedStage.
func (b *BuildManager) ShouldSkipStage(stageId constants.BuildStage) bool {
	for _, s := range b.skipStages {
		if s == stageId {
			return true
		}
	}
	return false
}

// Sets the active stage without the ability to go backwards.
func ToPreviousStage() {
	if (len(b.stack) >= 1) && (b.allowGoBack) {
		b.stack = b.stack[0 : len(b.stack)-1]
	}
}

// Sets active build stage for global access.
func SetActiveStage(stageId constants.BuildStage) {
	b.stack = append(b.stack, stageId)
}

// Sets active build stage for global access.
func SetAppConfigPath(p string) {
	b.AppConfigPath = p
}

// Sets active build stage for global access.
func SetConfigDirPath(p string) {
	b.ConfigDirPath = p
}

// Checks if build stage is active build stage.
func (b *BuildManager) IsActiveStage(stageId constants.BuildStage) bool {
	return b.Stage() == stageId
}

func (b *BuildManager) Stage() constants.BuildStage {
	return b.stack[len(b.stack)-1]
}

func (b *BuildManager) SendToPreviousStageMsg() tea.Cmd {
	ToPreviousStage()
	return signals.SetStage(b.stack[len(b.stack)-1])
}
