package cmd_option

import "github.com/spf13/cobra"

type CmdOptionBool struct {
	viperKey     string
	flagString   string
	shortHand    string
	defaultValue bool
	usageString  string
}

func (c CmdOptionBool) ViperKey() (_ string) {
	return c.viperKey
}

func (c CmdOptionBool) DefaultVal() (_ bool) {
	return c.defaultValue
}

func (c CmdOptionBool) FlagString() (_ string) {
	return c.flagString
}

func (c CmdOptionBool) Shorthand() (_ string) {
	return c.shortHand
}

func (c CmdOptionBool) Usage() (_ string) {
	return c.usageString
}

func (c CmdOptionBool) Init(cmd *cobra.Command) {
	if c.shortHand == "" {
		cmd.Flags().Bool(c.flagString, c.defaultValue, c.usageString)
	} else {
		cmd.Flags().BoolP(c.flagString, c.shortHand, c.defaultValue, c.usageString)
	}
}
