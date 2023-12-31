package cmd

import (
	"main/src/pkg/dagger"
	"main/src/pkg/zarf"

	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:     "package",
	Aliases: []string{"p"},
}

var packageCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "dpt package create",
	Long:  `dpt package create`,
	Run: func(cmd *cobra.Command, args []string) {
		runner := dagger.SetupDagger()
		defer runner.Client.Close()
		zarf.BuildZarfPackage(runner)
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)
	packageCmd.AddCommand(packageCreateCmd)
}
