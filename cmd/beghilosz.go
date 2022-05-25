package cmd

import (
	"strings"

	"github.com/kirillmorozov/encodor/beghilosz"
	"github.com/spf13/cobra"
)

func NewBeghiloszCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "beghilosz word...",
		Aliases: []string{"b"},
		Short:   "Turn input words into calculator spelling",
		Long: `Turn input words into calculator spelling.

Calculator spelling is an unintended characteristic of the seven-segments
display traditionally used by calculators, in which, when read upside-down, the
digits resemble letters of the Latin alphabet. Each digit may be mapped to one
or more letters, creating a limited but functional subset of the alphabet,
sometimes referred to as beghilos (or beghilosz).`,
		Example: "encodor beghilosz BOOBIES",
		Run: func(cmd *cobra.Command, args []string) {
			text := strings.Join(args, " ")
			encoded := beghilosz.Encode(text)
			cmd.Println(encoded)
		},
	}
}
