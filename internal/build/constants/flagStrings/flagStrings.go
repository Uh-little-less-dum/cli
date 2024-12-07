package flag_strings

type FlagString string

const (
	CloneTimeout  FlagString = "timeout"
	Here          FlagString = "here"
	AppConfigPath FlagString = "appConfig"
	LogFilePath   FlagString = "logFile"
)
