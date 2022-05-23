package cmd

import (
	"strings"

	"github.com/kirillmorozov/encodor/zalgo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(zalgoCmd)
}

var zalgoCmd = &cobra.Command{
	Use:     "zalgo word...",
	Aliases: []string{"z"},
	Short:   "Turn input words into criptic zalgo text",
	Long: `Turn input words into criptic zalgo text.
	
Zalgo text is digital text that has been modified with combining characters,
Unicode symbols used to add diacritics above or below letters, to appear
frightening or glitchy.`,
	Example: "encodor zalgo Ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn",
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		encoded := zalgo.Encode(text)
		cmd.Println(encoded)
	},
}
