package qrcrypt

import (
	"fmt"
	"github.com/LeuenbergerP/qrcrypt/pkg/qrcrypt"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Aliases: []string{"dec"},
	Short:   "Decrypt a string",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		keyPhrase := args[1]
		res := qrcrypt.Decrypt([]byte(args[0]), keyPhrase)
		fmt.Println(string(res))
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
