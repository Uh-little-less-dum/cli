package cmd_option

import (
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CmdOptionBool struct {
	ViperKey     viper_keys.ViperKey
	FlagString   string
	ShortHand    string
	DefaultValue bool
	UsageString  string
	EnvVar       string
}

func (c CmdOptionBool) Init(cmd *cobra.Command, v *viper.Viper) {
	v.SetDefault(string(c.ViperKey), c.DefaultValue)
	if c.ShortHand == "" {
		cmd.Flags().Bool(c.FlagString, c.DefaultValue, c.UsageString)
	} else {
		cmd.Flags().BoolP(c.FlagString, c.ShortHand, c.DefaultValue, c.UsageString)
	}

	err := v.BindPFlag(string(c.ViperKey), cmd.Flags().Lookup(c.FlagString))
	if err != nil {
		log.Fatal(err)
	}
	if c.EnvVar != "" {
		err = v.BindEnv(string(c.ViperKey), c.EnvVar)
		handleErr(err)
	}
}
