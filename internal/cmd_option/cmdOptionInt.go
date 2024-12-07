package cmd_option

import "github.com/spf13/cobra"

type CmdOptionInt struct {
	viperKey     string
	flagString   string
	shortHand    string
	defaultValue int
	usageString  string
}

func (c CmdOptionInt) ViperKey() (_ string) {
	return c.viperKey
}

func (c CmdOptionInt) DefaultVal() (_ int) {
	return c.defaultValue
}

func (c CmdOptionInt) FlagString() (_ string) {
	return c.flagString
}

func (c CmdOptionInt) Shorthand() (_ string) {
	return c.shortHand
}

func (c CmdOptionInt) Usage() (_ string) {
	return c.usageString
}

func (c CmdOptionInt) Init(cmd *cobra.Command) {
	if c.shortHand == "" {
		cmd.Flags().Int(c.flagString, c.defaultValue, c.usageString)
	} else {
		cmd.Flags().IntP(c.flagString, c.shortHand, c.defaultValue, c.usageString)
	}
}
