package buildConfig

import (
	"ulld/cli/internal/build/constants"

	"github.com/spf13/cobra"
)

type BuildConfigOpts struct {
	TargetDir    string
	InitialStage constants.BuildStage
}

// TODO: Set up the rest of the config flags here. Handle what you can with viper first though. Might not even need this struct if viper can handle everything.
func (c *BuildConfigOpts) Init(cmd *cobra.Command, targetDir string) {
	if len(targetDir) > 0 {
		c.TargetDir = targetDir
		c.InitialStage = constants.CloneTemplateAppStage
	} else {
		c.InitialStage = constants.ConfirmCurrentDirStage
	}
}
