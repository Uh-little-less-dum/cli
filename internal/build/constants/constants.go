package constants

import (
	tea "github.com/charmbracelet/bubbletea"
)

// FIXME: This is wrong. Need to find how to assign these types to the tea.msg type properly.
type ToRootModelMsg tea.Msg

type ConfirmDirectoryMsg tea.Msg

var ToRootModelCmd tea.Msg = tea.Msg("TRM")

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
)

const SparseCloneRepoUrl = "https://github.com/igloo1505/ulld.git"
const SparseCloneSparsePath = "apps/template"
