package cmd

import (
	"bytes"
	"strings"

	"github.com/kirillmorozov/encodor/zalgo"
	"github.com/spf13/cobra"
)

const (
	defaultDiacritics = 3
)

func NewZalgoCmd() *cobra.Command {
	zalgoCmd := &cobra.Command{
		Use:     "zalgo word...",
		Aliases: []string{"z"},
		Short:   "Turn input words into criptic zalgo text",
		Long: `Turn input words into criptic zalgo text.

Zalgo text is digital text that has been modified with combining characters,
Unicode symbols used to add diacritics above or below letters, to appear
frightening or glitchy.`,
		Example: "encodor zalgo -d 3 Ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn",
		Run: func(cmd *cobra.Command, args []string) {
			strength, flgErr := cmd.Flags().GetInt8("diacritics")
			if flgErr != nil {
				cmd.PrintErr(flgErr)
			}
			text := strings.Join(args, " ")
			input := strings.NewReader(text)
			output := bytes.NewBuffer(make([]byte, len(text)*int(strength)))
			encodingErr := zalgo.Encode(input, output, strength)
			if encodingErr != nil {
				cmd.PrintErr(encodingErr)
			}
			cmd.Println(output)
		},
	}
	zalgoCmd.Flags().Int8P(
		"diacritics",
		"d",
		defaultDiacritics,
		`How many diacritics are added to each letter.
Should be 1 <= diacritics <= 5`)
	return zalgoCmd
}
