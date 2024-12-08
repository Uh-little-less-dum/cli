package cli_config

import (
	"os"
	"path"

	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/Uh-little-less-dum/cli/internal/cmd_option"
	"github.com/charmbracelet/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CommandName int

const (
	RootCmdName CommandName = iota
	BuildCmdName
	MockCmdName
)

type ViperWrapper struct {
	viper            *viper.Viper
	defaultConfigDir string
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getConfigPaths(defaultCfgDir string) []string {
	additionalSources, ok := os.LookupEnv("ULLD_ADDITIONAL_SOURCES")

	if ok && (additionalSources != "") && (additionalSources != defaultCfgDir) {
		return []string{additionalSources, defaultCfgDir}
	} else {
		return []string{defaultCfgDir}
	}
}

// Set's viper.configDir property and appends all possible config paths to viper.
func (v *ViperWrapper) setConfigPaths(cmd *cobra.Command) {

	// TODO: Come back here and use the xdg package to set other possible paths, as well as to make sure that the default here is easily found by other programs.
	cfgPaths := getConfigPaths(v.defaultConfigDir)

	for _, c := range cfgPaths {
		v.viper.AddConfigPath(c)
	}

	v.viper.SetEnvPrefix("ULLD")
	v.viper.SetConfigName("cliConfig")
	v.viper.SetConfigType("json")

	cmd.PersistentFlags().StringP("configDir", "c", "", "config directory (default is $HOME/.ulld or $ULLD_ADDITIONAL_SOURCES)")
	err := v.viper.BindPFlag("configDir", cmd.PersistentFlags().Lookup("configDir"))
	handleErr(err)
	err = v.viper.BindEnv("configDir", "ULLD_ADDITIONAL_SOURCES")
	handleErr(err)
}

// Reads appConfig from file, and if it exists sets that directory to the viper.configDir variable.
func (v *ViperWrapper) readConfig() {
	if err := v.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug("No cliConfig.json file found. Continuing anyways.")
		} else {
			// Config file was found but another error was produced
			log.Error(err)
			os.Exit(1)
		}
	}
	cfgFile := v.viper.ConfigFileUsed()
	if cfgFile != "" {
		dirPath := path.Dir(cfgFile)
		if dirPath != "" {
			v.viper.Set(string(viper_keys.ConfigDir), dirPath)
		}
	}
}

func (v *ViperWrapper) setConfigDefaults() {
	home, err := homedir.Dir()
	if err != nil {
		log.Error("Error occurred while locating home directory.")
		os.Exit(1)
	}
	defaultCfgDir := path.Join(home, ".ulld")

	v.viper.SetDefault("configDir", defaultCfgDir)
	v.defaultConfigDir = defaultCfgDir
}

// Sets all cobra flags and binds them to viper.
func (v *ViperWrapper) setFlags(cmd *cobra.Command, cmdOptions []cmd_option.CmdOption) {
	for _, o := range cmdOptions {
		o.Init(cmd, v.viper)
	}
}

func (v *ViperWrapper) applyLogLevel() {

	// llString := cmp.Or(v.viper.GetString(string(viper_keys.LogLevel)), "info")

	llString := v.viper.GetString(string(viper_keys.LogLevel))

	if llString != "" {
		parsedLevel, err := log.ParseLevel(llString)
		if err != nil {
			log.Debugf("Provided log level of %s is not a supported log level.", llString)
		} else {
			log.SetLevel(parsedLevel)
		}
	}
}

func (v *ViperWrapper) Init(cmd *cobra.Command, opts []cmd_option.CmdOption) {
	v.viper = viper.GetViper()

	v.setConfigDefaults()
	v.setConfigPaths(cmd)

	v.viper.AutomaticEnv()
	v.setFlags(cmd, opts)

	v.readConfig()

	v.applyLogLevel()
}

func InitViper(cmd *cobra.Command, cmdOpts []cmd_option.CmdOption) func() {
	return func() {
		var v = ViperWrapper{}
		v.Init(cmd, cmdOpts)
	}
}
