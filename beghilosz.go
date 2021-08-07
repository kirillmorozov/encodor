package encodor

import (
	"strings"
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
		words = reverseSlice(words)
		lines[lineIndex] = strings.Join(words, " ")
	}
	lines = reverseSlice(lines)
	return strings.Join(lines, "\n")
}

func isSpecialWord(word string) bool {
	isHashtag := strings.HasPrefix(word, "#")
	isUsername := strings.HasPrefix(word, "@")
	return isHashtag || isUsername
}

func reverseString(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func reverseSlice(slice []string) []string {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
}
