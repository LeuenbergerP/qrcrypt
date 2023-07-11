package qrcrypt

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/LeuenbergerP/qrcrypt/pkg/qrcrypt"
	"github.com/mdp/qrterminal/v3"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Aliases: []string{"enc"},
	Short:   "Encrypt a string",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		content := args[1]
		keyPhrase := args[2]
		res := qrcrypt.Encrypt([]byte(content), keyPhrase)
		writeToFile(res, fileName)
	},
}

func writeToFile(data []byte, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()
	encoded := base64Encode(data)
	qrCode := qrCode(encoded)
	fmt.Println(qrCode)
	fmt.Println(encoded)
	_, err2 := f.WriteString(encoded)
	if err2 != nil {
		log.Fatal(err2)
		return err2
	}
	return nil
}

func base64Encode(data []byte) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}

func qrCode(val string) string {
	buf := new(bytes.Buffer)
	config := qrterminal.Config{
		Level:      qrterminal.L,
		Writer:     buf,
		HalfBlocks: false,
		BlackChar:  qrterminal.WHITE,
		WhiteChar:  qrterminal.BLACK,
		QuietZone:  4,
	}
	qrterminal.GenerateWithConfig(val, config)
	return buf.String()
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
