package build_stage_utils_test

import (
	"os"
	"testing"

	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	build_stage_utils "github.com/Uh-little-less-dum/cli/internal/buildStageManagement"
	"github.com/Uh-little-less-dum/cli/internal/mocks"
	cli_config "github.com/Uh-little-less-dum/cli/internal/utils/initViper"
	// "github.com/Uh-little-less-dum/build/internal/build/constants"
	// stage_gather_config_location "github.com/Uh-little-less-dum/build/internal/buildScript/stages/gather_config_location"
	// "github.com/Uh-little-less-dum/build/internal/mocks"
	// cli_config "github.com/Uh-little-less-dum/build/internal/utils/initViper"
)

func Test_GetNextBuildStage(t *testing.T) {
	originalAdditionalSource := os.Getenv("ULLD_ADDITIONAL_SOURCES")
	mocks.MockCommandSetup(cli_config.BuildCmdName)
	t.Run("Finds file according to environment variable", func(t *testing.T) {
		configPath, stage := build_stage_utils.GetNextBuildStage()
		if (configPath == "") || (stage != constants.ConfirmConfigLocFromEnv) {
			t.Fail()
		}
	})

	t.Run("Returns as if not found with env variable but no file", func(t *testing.T) {
		os.Setenv("ULLD_ADDITIONAL_SOURCES", "~/Desktop/test")
		configPath, stage := build_stage_utils.GetNextBuildStage()
		if (configPath != "") || (stage != constants.ChooseWaitOrPickConfigLoc) {
			t.Fail()
		}
	})

	t.Run("Returns as if not found with no env variable", func(t *testing.T) {
		os.Setenv("ULLD_ADDITIONAL_SOURCES", "")
		configPath, stage := build_stage_utils.GetNextBuildStage()
		if (configPath != "") || (stage != constants.ChooseWaitOrPickConfigLoc) {
			t.Fail()
		}
	})
	os.Setenv("ULLD_ADDITIONAL_SOURCES", originalAdditionalSource)
}
