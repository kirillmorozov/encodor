// Package beghilosz provides Encode function that encodes text using calculator
// spelling.
//
// Calculator spelling is an unintended characteristic of the seven-segments
// display traditionally used by calculators, in which, when read upside-down,
// the digits resemble letters of the Latin alphabet. Each digit may be mapped
// to one or more letters, creating a limited but functional subset of the
// alphabet, sometimes referred to as beghilos (or beghilosz).
package beghilosz

import (
	"strings"

	"github.com/kirillmorozov/encodor/utils"
)

var beghiloszReplacer = strings.NewReplacer(
	"B", "8",
	"E", "3",
	"G", "6",
	"H", "4",
	"I", "1",
	"L", "7",
	"O", "0",
	"S", "5",
	"Z", "2",
)

// Encode encodes text into calculator spelling.
//
// Special words(as determined by utils.IsSpecialWord) are
// left as is.
func Encode(text string) string {
	text = strings.ToUpper(text)
	words := strings.Split(text, " ")
	var builder strings.Builder
	builder.Grow(len(text))
	for i := len(words) - 1; i >= 0; i-- {
		if !utils.IsSpecialWord(words[i]) {
			words[i] = beghiloszReplacer.Replace(words[i])
			words[i] = utils.ReverseString(words[i])
		}
		builder.WriteString(words[i])
		if i != 0 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}
