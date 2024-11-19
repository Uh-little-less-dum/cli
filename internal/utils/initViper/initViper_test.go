package cli_config

import (
	"fmt"
	"os"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var TestCmd = &cobra.Command{
	Use:   "dumTest",
	Short: "",
	Long:  ``,
}

func Test_InitViperLogLevel(t *testing.T) {
	globalLogLevel, ok := os.LookupEnv("ULLD_LOG_LEVEL")
	if !ok {
		os.Setenv("ULLD_LOG_LEVEL", "debug")
		globalLogLevel = "debug"
	}
	var vals = []struct {
		name     string
		inputVal string
		// expected string
	}{
		{"logLevel set properly from environment", globalLogLevel},
		{"logLevel set properly from environment", "debug"},
		{"logLevel set properly from environment", "info"},
		{"logLevel set properly from environment", "warn"},
	}
	cmd := TestCmd
	InitViper(cmd, BuildCmdName)()
	for _, tt := range vals {
		os.Setenv("ULLD_LOG_LEVEL", tt.inputVal)
		t.Run(tt.name, func(t *testing.T) {
			expected, err := log.ParseLevel(tt.inputVal)
			if err != nil {
				t.Fatal(err)
			}
			value := viper.GetViper().GetString("logLevel")
			es := expected.String()
			if es == "" {
				t.Error("Expected returned an empty string")
			}
			if value != es {
				t.Errorf("Expected '%s', received '%s'", es, value)
			}
		})
	}
}

func Test_InitViperAdditionalSources(t *testing.T) {
	globalLogLevel, ok := os.LookupEnv("ULLD_ADDITIONAL_SOURCES")
	if !ok {
		os.Setenv("ULLD_LOG_LEVEL", "debug")
		globalLogLevel = "debug"
	}
	var vals = []struct {
		name     string
		inputVal string
	}{
		{"ULLD_ADDITIONAL_SOURCES set from environment", globalLogLevel},
		{"ULLD_ADDITIONAL_SOURCES set from environment", "~/dev-utils/ulld/"},
	}
	cmd := TestCmd
	for _, tt := range vals {
		cmd.ResetFlags()
		viper.Reset()
		os.Setenv("ULLD_ADDITIONAL_SOURCES", tt.inputVal)
		InitViper(cmd, BuildCmdName)()
		t.Run(tt.name, func(t *testing.T) {
			value := viper.GetViper().GetString("configDir")
			if value != tt.inputVal {
				t.Errorf("Expected '%s', received '%s'", tt.inputVal, value)
			}
		})
	}
}

type TestItem struct {
	name     string
	viperKey string
	flagKey  string
	inputVal string
}

func Test_Flags(t *testing.T) {
	var vals = []TestItem{
		{"logFile", "logFile", "logFile", "/Users/bigsexy/Desktop/current/ulld/buildUtils/testLog.log"},
		{"timeout", "timeout", "timeout", "30"},
	}
	cmd := TestCmd
	for _, tt := range vals {
		viper.Reset()
		cmd.ResetFlags()
		InitViper(cmd, BuildCmdName)()
		err := cmd.Flags().Set(tt.flagKey, tt.inputVal)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			val := viper.GetViper().GetString(tt.viperKey)
			if val != tt.inputVal {
				t.Errorf("Expected %s, Received %s", tt.inputVal, val)
			}
		})
	}
}

func Test_ViperDefaults(t *testing.T) {
	var vals = []struct {
		viperKey string
		expected any
	}{
		{"timeout", 30},
	}
	cmd := TestCmd
	for _, tt := range vals {
		viper.Reset()
		cmd.ResetFlags()
		InitViper(cmd, BuildCmdName)()
		t.Run(fmt.Sprintf("Default value for %s", tt.viperKey), func(t *testing.T) {
			val := viper.GetViper().Get(tt.viperKey)
			if val != tt.expected {
				t.Errorf("Expected %s, Received %v", tt.expected, val)
			}
		})
	}
}
