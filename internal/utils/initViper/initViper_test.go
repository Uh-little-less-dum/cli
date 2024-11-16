package cli_config

import (
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	InitViper(cmd)()
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
	InitViper(cmd)()
	for _, tt := range vals {
		os.Setenv("ULLD_ADDITIONAL_SOURCES", tt.inputVal)
		t.Run(tt.name, func(t *testing.T) {
			value := viper.GetViper().GetString("configDir")
			if value != tt.inputVal {
				t.Errorf("Expected '%s', received '%s'", tt.inputVal, value)
			}
		})
	}
}

func Test_Flags(t *testing.T) {
	var vals = []struct {
		name     string
		viperKey string
		flagKey  string
		inputVal string
	}{
		{"logFile", "logFile", "logFile", "/Users/bigsexy/Desktop/current/ulld/buildUtils/testLog.log"},
	}
	cmd := TestCmd
	for _, tt := range vals {
		viper.Reset()
		InitViper(cmd)()
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
