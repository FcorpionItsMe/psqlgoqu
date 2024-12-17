package cmd

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/cmd/mktab"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goqu",
	Short: "GOQU is the super library for making queries for postgresql.",
	Long: `GOQU is the super library for making queries for postgresql.
                It also generate models for queries!`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	InitCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func InitCommands() {
	mktab.InitCommand(rootCmd)
}
