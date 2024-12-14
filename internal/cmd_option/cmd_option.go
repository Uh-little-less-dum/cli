package cmd_option

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CmdOption interface {
	Init(cmd *cobra.Command, v *viper.Viper)
}
