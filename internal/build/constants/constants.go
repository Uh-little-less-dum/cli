package constants

type BuildStage int

const (
	ConfirmCurrentDirStage BuildStage = iota
	PickTargetDirStage
	ConfirmConfigLocFromEnv
	PickConfigLoc
	ConfirmWaitForConfigMove
	WaitForConfigMove
	ChooseWaitOrPickConfigLoc
	CloneTemplateAppStage
	InstallTemplateAppDeps
	PreConflictResolveBuildStream
)
