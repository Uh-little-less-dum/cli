package build_config_test

import (
	"testing"

	build_config "github.com/Uh-little-less-dum/cli/internal/build/config"
	"github.com/Uh-little-less-dum/cli/internal/build/constants"
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/spf13/viper"
)

type ViperItem struct {
	viperKey viper_keys.ViperKey
	value    any
}

func Test_BuildConfigInitStage(t *testing.T) {
	var vals = []struct {
		name       string
		args       []string
		viperItems []ViperItem
		expected   constants.BuildStage
	}{
		{"With target dir and not useCwd flag or appConfig path.", []string{"/Users/bigsexy/Desktop"}, []ViperItem{}, constants.ConfirmConfigLocFromEnv},
		{"With target dir and appConfig path.", []string{"/Users/bigsexy/Desktop"}, []ViperItem{
			{viper_keys.AppConfigPath, "/Users/bigsexy/dev-utils/ulld/appConfig.ulld.json"},
		}, constants.CloneTemplateAppStage},
		{"Without target dir or appConfig path.", []string{}, []ViperItem{}, constants.ConfirmCurrentDirStage},
	}
	for _, tt := range vals {
		viper.Reset()
		for _, vi := range tt.viperItems {
			viper.Set(string(vi.viperKey), vi.value)
		}
		t.Run(tt.name, func(t *testing.T) {
			b := build_config.GetBuildManager()
			b.Init(tt.args)
			s := b.Stage()
			if s != tt.expected {
				t.Errorf("Expected '%d', received '%d'", int(tt.expected), int(s))
			}
		})
	}
}
