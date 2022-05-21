package cmd

import (
	"fmt"
	"strings"

	"github.com/kirillmorozov/encodor/beghilosz"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(beghiloszCmd)
}

var beghiloszCmd = &cobra.Command{
	Use:   "beghilosz",
	Short: "Encode text using calculator spelling",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		encoded := beghilosz.Encode(text)
		fmt.Println(encoded)
	},
}
