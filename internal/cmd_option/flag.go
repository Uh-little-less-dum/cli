package cmd_option

import (
	"github.com/spf13/cobra"
)

type CmdOption[T string | int | bool | float32] interface {
	ViperKey() string
	FlagString() string
	Shorthand() string
	Usage() string
	Init(cmd *cobra.Command)
	DefaultVal() T
}
