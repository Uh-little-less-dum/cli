package viper_keys

type ViperKey string

const (
	// Directory where ULLD will be generated.
	TargetDirectory ViperKey = "targetDir"
	// Timeout as it pertains to the git clone and npm/pnpm/yarn install requests.
	CloneTimeout ViperKey = "cloneTimeout"
	UseCwd       ViperKey = "useCwd"
	// Path to appConfig.ulld.json or equivalent file to be used.
	AppConfigPath ViperKey = "appConfig"
	LogFilePath   ViperKey = "logFile"
	LogLevel      ViperKey = "logLevel"
	// Path to the ULLD_ADDITIONAL_SOURCES directory if one exists.
	ConfigDir ViperKey = "configDir"
	// Which package manager to use. "pnpm" | "npm" | "yarn"
	PackageManager ViperKey = "packageManager"
)
