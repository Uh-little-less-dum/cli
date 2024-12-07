package cmd_option

import "github.com/spf13/cobra"

type CmdOptionFloat struct {
	viperKey     string
	flagString   string
	shortHand    string
	defaultValue float32
	usageString  string
}

func (c CmdOptionFloat) ViperKey() (_ string) {
	return c.viperKey
}

func (c CmdOptionFloat) DefaultVal() (_ float32) {
	return c.defaultValue
}

func (c CmdOptionFloat) FlagString() (_ string) {
	return c.flagString
}

func (c CmdOptionFloat) Shorthand() (_ string) {
	return c.shortHand
}

func (c CmdOptionFloat) Usage() (_ string) {
	return c.usageString
}

func (c CmdOptionFloat) Init(cmd *cobra.Command) {
	if c.shortHand == "" {
		cmd.Flags().Float32(c.flagString, c.defaultValue, c.usageString)
	} else {
		cmd.Flags().Float32P(c.flagString, c.shortHand, c.defaultValue, c.usageString)
	}
}
