package qrcrypt

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/LeuenbergerP/qrcrypt/pkg/qrcrypt"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Aliases: []string{"dec"},
	Short:   "Decrypt a string",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		keyPhrase := args[1]
		content, err := readFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res := qrcrypt.Decrypt(content, keyPhrase)
		fmt.Println(string(res))
	},
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, fi.Size())
	_, err2 := f.Read(data)
	if err2 != nil {
		return nil, err2
	}
	b64Decoded, err3 := base64Decode(data)
	if err3 != nil {
		return nil, err3
	}
	return b64Decoded, nil
}

func base64Decode(data []byte) ([]byte, error) {
	return b64.StdEncoding.DecodeString(string(data))
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
