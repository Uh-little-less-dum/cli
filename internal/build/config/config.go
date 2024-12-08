package build_config

import (
	"os"

	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	"github.com/charmbracelet/log"
)

type BuildManager struct {
	TargetDir    string
	currentStage constants.BuildStage
	initialStage constants.BuildStage
	skipStages   []constants.BuildStage
}

var b *BuildManager

// Returns the BuildManager singleton.
func GetBuildManager() *BuildManager {
	return b
}

func init() {
	val := BuildManager{}
	if val.TargetDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		val.TargetDir = cwd
	}
	b = &val
}

func (b *BuildManager) InitialStage() constants.BuildStage {
	return b.initialStage
}

func (b *BuildManager) SetInitialStage(s constants.BuildStage) {
	b.initialStage = s
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
func (c *BuildManager) Init() {
	c.SetInitialStage(constants.ConfirmCurrentDirStage)
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

// Sets active build stage for global access.
func SetActiveStage(stageId constants.BuildStage) {
	b.currentStage = stageId
}

// Checks if build stage is active build stage.
func (b *BuildManager) IsActiveStage(stageId constants.BuildStage) bool {
	return b.currentStage == stageId
}
