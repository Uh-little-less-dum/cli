package build_config

import (
	"github.com/Uh-little-less-dum/cli/internal/build/constants"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type BuildConfigOpts struct {
	targetDir    string
	initialStage constants.BuildStage
}

const skipStagesKey string = "skippedBuildStages"
const activeStageKey string = "activeBuildStage"

func (b BuildConfigOpts) TargetDir() string {
	return b.targetDir
}

func (b *BuildConfigOpts) SetTargetDir(newDir string) {
	b.targetDir = newDir
}

func (b *BuildConfigOpts) InitialStage() constants.BuildStage {
	return b.initialStage
}

func (b *BuildConfigOpts) SetInitialStage(s constants.BuildStage) {
	b.initialStage = s
}

// Adds a build stage that should be skipped to the list of skipped stages. This list can then be accessed from within each build stage at which point the stage can be bypassed.
func (b *BuildConfigOpts) AddSkippedStage(stage constants.BuildStage) {
	v := viper.GetViper()
	currentlySkippedStages := v.GetIntSlice(skipStagesKey)
	v.Set(skipStagesKey, append(currentlySkippedStages, int(stage)))
}

// TODO: Set up the rest of the config flags here. Handle what you can with viper first though. Might not even need this struct if viper can handle everything.
func (c *BuildConfigOpts) Init(cmd *cobra.Command) {
	c.SetInitialStage(constants.ConfirmCurrentDirStage)
}

// Utility to check if stage should be skipped. Reads data formerly set by AddSkippedStage.
func ShouldSkipStage(stageId constants.BuildStage) bool {
	skipStages := viper.GetViper().GetIntSlice(skipStagesKey)
	stage := int(stageId)
	for _, s := range skipStages {
		if s == stage {
			return true
		}
	}
	return false
}

// Sets active build stage for global access.
func SetActiveStage(stageId constants.BuildStage) {
	viper.GetViper().Set(activeStageKey, stageId)
}

// Checks if build stage is active build stage.
func IsActiveStage(stageId constants.BuildStage) bool {
	val := viper.GetViper().GetInt(activeStageKey)
	return val == int(stageId)
}
