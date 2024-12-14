package stage_gather_config_location_test

import (
	"os"
	"testing"

	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	stage_gather_config_location "github.com/Uh-little-less-dum/cli/internal/buildScript/stages/gather_config_location"
	"github.com/Uh-little-less-dum/cli/internal/mocks"
	_ "github.com/Uh-little-less-dum/cli/internal/mocks"
	cli_config "github.com/Uh-little-less-dum/cli/internal/utils/initViper"
)

func Test_GetNextBuildStage(t *testing.T) {
	originalAdditionalSource := os.Getenv("ULLD_ADDITIONAL_SOURCES")
	mocks.MockCommandSetup(cli_config.BuildCmdName)
	t.Run("Finds file according to environment variable", func(t *testing.T) {
		configPath, stage := stage_gather_config_location.GetNextBuildStage()
		if (configPath == "") || (stage != build_stages.ConfirmConfigLocFromEnv) {
			t.Fail()
		}
	})

	t.Run("Returns as if not found with env variable but no file", func(t *testing.T) {
		os.Setenv("ULLD_ADDITIONAL_SOURCES", "~/Desktop/test")
		configPath, stage := stage_gather_config_location.GetNextBuildStage()
		if (configPath != "") || (stage != build_stages.ChooseWaitOrPickConfigLoc) {
			t.Fail()
		}
	})

	t.Run("Returns as if not found with no env variable", func(t *testing.T) {
		os.Setenv("ULLD_ADDITIONAL_SOURCES", "")
		configPath, stage := stage_gather_config_location.GetNextBuildStage()
		if (configPath != "") || (stage != build_stages.ChooseWaitOrPickConfigLoc) {
			t.Fail()
		}
	})
	os.Setenv("ULLD_ADDITIONAL_SOURCES", originalAdditionalSource)
}
