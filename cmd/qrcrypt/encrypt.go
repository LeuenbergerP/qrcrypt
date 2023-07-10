package qrcrypt

import (
	"fmt"
	"github.com/LeuenbergerP/qrcrypt/pkg/qrcrypt"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Aliases: []string{"enc"},
	Short:   "Encrypt a string",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		keyPhrase := args[1]
		res := qrcrypt.Encrypt([]byte(args[0]), keyPhrase)
		fmt.Println(string(res))
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
