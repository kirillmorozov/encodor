package cmd

import (
	"fmt"
	"strings"

	"github.com/kirillmorozov/encodor/zalgo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(zalgoCmd)
}

var zalgoCmd = &cobra.Command{
	Use:   "zalgo",
	Short: "Encode text using criptic zalgo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		encoded := zalgo.Encode(text)
		fmt.Println(encoded)
	},
}
