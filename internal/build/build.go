package build

import (
	"ulld/cli/internal/build/constants"
	mainBuildModel "ulld/cli/internal/build/ui/mainmodel"
)

// func StartBuildUI() {

// }

func BuildUlld(e UlldEnv) {
	// TODO: Handle this initial stage based on passed flags.
	mm := mainBuildModel.InitialMainModel(constants.PickTargetDirStage)
	// targetDir := SelectTargetDir()
	// logger.DebugLog(targetDir)
}
