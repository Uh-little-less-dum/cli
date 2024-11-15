package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func Test_RootCmdEnvVars(t *testing.T) {
	additionalSources, hasVal := os.LookupEnv("ULLD_ADDITIONAL_SOURCES")
	if !hasVal {
		t.Fatal("No ULLD_ADDITIONAL_SOURCES variable found.")
	}
	cmd := RootCmd.Root()
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)
	var tests = []struct {
		name      string
		envKey    string
		viperKey  string
		testValue string
	}{
		{"configDir", "ULLD_ADDITIONAL_SOURCES", "configDir", additionalSources},
		{"logLevel", "ULLD_LOG_LEVEL", "logLevel", "debug"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.envKey, tt.testValue)
			err := cmd.Execute()
			if err != nil {
				t.Error(err)
			}
			viperVar := viper.Get(tt.viperKey)
			if viperVar != tt.testValue {
				t.Errorf("EnvKey %s != ViperKey %s", tt.envKey, tt.viperKey)
			}
		})
	}
}
