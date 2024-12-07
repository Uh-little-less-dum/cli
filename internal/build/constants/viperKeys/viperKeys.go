package viper_keys

type ViperKey string

const (
	TargetDirectory ViperKey = "targetDir"
	CloneTimeout    ViperKey = "cloneTimeout"
	UseCwd          ViperKey = "useCwd"
	AppConfigPath   ViperKey = "appConfig"
	LogFilePath     ViperKey = "logFile"
)
