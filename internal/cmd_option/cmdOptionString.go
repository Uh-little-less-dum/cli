package cmd_option

import "github.com/spf13/cobra"

type CmdOptionString struct {
	viperKey     string
	flagString   string
	shortHand    string
	defaultValue string
	usageString  string
}

func (c CmdOptionString) ViperKey() (_ string) {
	return c.viperKey
}

func (c CmdOptionString) DefaultVal() (_ string) {
	return c.defaultValue
}

func (c CmdOptionString) FlagString() (_ string) {
	return c.flagString
}

func (c CmdOptionString) Shorthand() (_ string) {
	return c.shortHand
}

func (c CmdOptionString) Usage() (_ string) {
	return c.usageString
}

// RESUME: Come back here and implement this Init field in each of the different CmdOption models, then set all available flags as an array of CmdOption models before reworking the build init methods to generate cobra stuff from that list alone.
func (c CmdOptionString) Init(cmd *cobra.Command) {
	if c.shortHand == "" {
		cmd.Flags().String(c.flagString, c.defaultValue, c.usageString)
	} else {
		cmd.Flags().StringP(c.flagString, c.shortHand, c.defaultValue, c.usageString)
	}
}
