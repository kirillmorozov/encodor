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

// Encode transforms text into calculator spelling.
//
// Hashtags(words beginning with `#`) and mentions(words beginning with `@`) are
// left as is.
func Encode(text string) string {
	beghiloszReplacer := strings.NewReplacer(
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
	text = strings.ToUpper(text)
	lines := strings.Split(text, "\n")
	for lineIndex, line := range lines {
		words := strings.Fields(line)
		for wordIndex, word := range words {
			if !utils.IsSpecialWord(word) {
				word = beghiloszReplacer.Replace(word)
				word = utils.ReverseString(word)
			}
			words[wordIndex] = word
		}
		words = utils.ReverseStringSlice(words)
		lines[lineIndex] = strings.Join(words, " ")
	}
	lines = utils.ReverseStringSlice(lines)
	return strings.Join(lines, "\n")
}
