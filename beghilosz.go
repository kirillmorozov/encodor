package encodor

import (
	"strings"
)

// beghiloszReplacer replace all runes according to beghilosz encoding.
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

// Beghilosz encode text to calculator spelling.
// Hashtags(words beginning with '#') and mentions(words beginning with '@') are
// left as is.
func Beghilosz(text string) string {
	text = strings.ToUpper(text)
	lines := strings.Split(text, "\n")
	for lineIndex, line := range lines {
		words := strings.Fields(line)
		for wordIndex, word := range words {
			if !isSpecialWord(word) {
				word = beghiloszReplacer.Replace(word)
				word = reverseString(word)
			}
			words[wordIndex] = word
		}
		words = reverseStringSlice(words)
		lines[lineIndex] = strings.Join(words, " ")
	}
	lines = reverseStringSlice(lines)
	return strings.Join(lines, "\n")
}

// isSpecialWord return true if word is a hashtag or a username.
func isSpecialWord(word string) bool {
	isHashtag := strings.HasPrefix(word, "#")
	isUsername := strings.HasPrefix(word, "@")
	return isHashtag || isUsername
}

// reverseString reverse rune order in str.
func reverseString(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// reverseStringSlice reverse strings order in slice.
func reverseStringSlice(slice []string) []string {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
}
