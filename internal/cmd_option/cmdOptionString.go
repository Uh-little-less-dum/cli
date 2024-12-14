package cmd_option

import (
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CmdOptionString struct {
	ViperKey     viper_keys.ViperKey
	FlagString   string
	ShortHand    string
	DefaultValue string
	UsageString  string
	EnvVar       string
}

func (c CmdOptionString) Init(cmd *cobra.Command, v *viper.Viper) {
	if c.EnvVar != "" {
		err := v.BindEnv(string(c.ViperKey), c.EnvVar)
		handleErr(err)
	}
	if c.DefaultValue != "" {
		v.SetDefault(string(c.ViperKey), c.DefaultValue)
	}
	if c.ShortHand == "" {
		cmd.Flags().String(c.FlagString, c.DefaultValue, c.UsageString)
	} else {
		cmd.Flags().StringP(c.FlagString, c.ShortHand, c.DefaultValue, c.UsageString)
	}

	err := v.BindPFlag(string(c.ViperKey), cmd.Flags().Lookup(c.FlagString))
	if err != nil {
		log.Fatal(err)
	}
}
