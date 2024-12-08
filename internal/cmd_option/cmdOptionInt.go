package cmd_option

import (
	viper_keys "github.com/Uh-little-less-dum/cli/internal/build/constants/viperKeys"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CmdOptionInt struct {
	ViperKey     viper_keys.ViperKey
	FlagString   string
	ShortHand    string
	DefaultValue int
	UsageString  string
	EnvVar       string
}

func (c CmdOptionInt) Init(cmd *cobra.Command, v *viper.Viper) {
	if c.DefaultValue != 0 {
		v.SetDefault(string(c.ViperKey), c.DefaultValue)
	}
	if c.ShortHand == "" {
		cmd.Flags().Int(c.FlagString, c.DefaultValue, c.UsageString)
	} else {
		cmd.Flags().IntP(c.FlagString, c.ShortHand, c.DefaultValue, c.UsageString)
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
