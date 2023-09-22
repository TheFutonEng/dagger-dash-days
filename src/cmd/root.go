package cmd

import (
	"github.com/spf13/cobra"
)

var (
	theCmd = "dpt"
)

var rootCmd = &cobra.Command{
	Use:   theCmd + " COMMAND",
	Short: "Run " + theCmd + " COMMAND --help for more information on a command.",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
