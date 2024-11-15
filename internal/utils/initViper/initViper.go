package cli_config

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	var cfgDir string

	cmd.PersistentFlags().StringVarP(&cfgDir, "configDir", "c", "", "config directory (default is $HOME/.ulld or $ULLD_ADDITIONAL_SOURCES)")
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
	dirPath := path.Dir(cfgFile)
	v.viper.Set("configDir", dirPath)
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
func (v *ViperWrapper) setFlags(cmd *cobra.Command) {
	// Log File
	cmd.Flags().StringP("logFile", "l", "", "Log output to this file. Useful for build failures and other local development.")
	err := v.viper.BindPFlag("logFile", cmd.Flags().Lookup("logFile"))
	handleErr(err)

	// Log Level
	cmd.Flags().String("logLevel", "info", "Log level")
	err = v.viper.BindEnv("logLevel", "ULLD_LOG_LEVEL")
	handleErr(err)
	err = v.viper.BindPFlag("logLevel", cmd.Flags().Lookup("logLevel"))
	handleErr(err)
}

func (v *ViperWrapper) setLogLevel(cmd *cobra.Command) {

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

	v.setLogLevel(cmd)

}

func InitViper(cmd *cobra.Command) func() {
	return func() {
		var v = ViperWrapper{}
		v.Init(cmd)
	}
}
