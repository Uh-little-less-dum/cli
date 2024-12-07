package cli_config

import (
	"os"
	"path"

	flag_strings "github.com/Uh-little-less-dum/cli/internal/build/constants/flagStrings"
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/Uh-little-less-dum/cli/internal/flag"
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
			v.viper.Set("configDir", dirPath)
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
func (v *ViperWrapper) setFlags(cmd *cobra.Command, cmdOptions []flag.CmdOption) {
	for _, o := range cmdOptions {
		o.Init(cmd)
	}
	// Log File
	cmd.Flags().StringP(string(flag_strings.LogFilePath), "l", "", "Log output to this file. Useful for build failures and other local development.")
	err := v.viper.BindPFlag(string(viper_keys.LogFilePath), cmd.Flags().Lookup(string(flag_strings.LogFilePath)))
	handleErr(err)

	// Log Level
	cmd.Flags().String("logLevel", "info", "Log level")
	err = v.viper.BindEnv("logLevel", "ULLD_LOG_LEVEL")
	handleErr(err)
	err = v.viper.BindPFlag("logLevel", cmd.Flags().Lookup("logLevel"))
	handleErr(err)
}

func (v *ViperWrapper) applyLogLevel() {

	llString := v.viper.GetString("logLevel")

	if llString != "" {
		parsedLevel, err := log.ParseLevel(llString)
		if err != nil {
			log.Debugf("Provided log level of %s is not a supported log level.", llString)
		} else {
			log.SetLevel(parsedLevel)
		}
	}
}

func (v *ViperWrapper) Init(cmd *cobra.Command) {
	v.viper = viper.GetViper()

	v.setConfigDefaults()
	v.setConfigPaths(cmd)

	v.viper.AutomaticEnv()
	v.setFlags(cmd)

	v.readConfig()

	v.applyLogLevel()
}

func (v *ViperWrapper) InitBuildCmd(cmd *cobra.Command) {
	// Timeout flag
	v.viper.SetDefault(string(viper_keys.CloneTimeout), 30)
	cmd.Flags().Int(string(flag_strings.CloneTimeout), 30, "Log level")
	err := v.viper.BindPFlag(string(viper_keys.CloneTimeout), cmd.Flags().Lookup(string(flag_strings.CloneTimeout)))
	handleErr(err)

	// Bypass location select flag
	cmd.Flags().Bool(string(flag_strings.Here), false, "Bypass the directory selection input and use the current working directory.")
	err = v.viper.BindPFlag(string(viper_keys.UseCwd), cmd.Flags().Lookup(string(flag_strings.Here)))
	handleErr(err)

	// appConfig.ulld.json path.
	cmd.Flags().StringP(string(flag_strings.AppConfigPath), "a", "", "Bypass the file path select menu and use this path for your appConfig.ulld.json source.")
	err = v.viper.BindPFlag(string(viper_keys.AppConfigPath), cmd.Flags().Lookup(string(flag_strings.AppConfigPath)))
	handleErr(err)
}

func InitViper(cmd *cobra.Command, buildName CommandName) func() {
	return func() {
		var v = ViperWrapper{}
		v.Init(cmd)
		switch buildName {
		case BuildCmdName:
			v.InitBuildCmd(cmd)
		}
	}
}
