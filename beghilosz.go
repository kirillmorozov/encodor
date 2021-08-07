package encodor

import (
	"strings"
)

func Beghilosz(input string) string {
	lines := strings.Split(input, "\n")
	replacer := strings.NewReplacer(
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
	for lineIndex, line := range lines {
		words := strings.Fields(line)
		for wordIndex, word := range words {
			word = strings.ToUpper(word)
			if !isSpecialWord(word) {
				word = replacer.Replace(word)
				word = reverseString(word)
			}
			words[wordIndex] = word
		}
		words = reverseSlice(words)
		lines[lineIndex] = strings.Join(words, " ")
	}
	return strings.Join(reverseSlice(lines), "\n")
}

func isSpecialWord(word string) bool {
	isHashtag := strings.HasPrefix(word, "#")
	isUsername := strings.HasPrefix(word, "@")
	return isHashtag || isUsername
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func reverseSlice(input []string) []string {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
	return input
}
