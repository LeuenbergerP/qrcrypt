package qrcrypt

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "qrcrypt",
	Short: "qrcrypt - a simple CLI to transform and inspect strings into sha encrypted QR codes",
	Long:  `qrcrypt is a simple CLI to transform and inspect strings into sha encrypted QR codes.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
