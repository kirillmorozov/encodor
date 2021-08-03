package encodor

import (
	"strings"
)

func beghiloszMapping(letter rune) rune {
	var beghiloszRuneMap = map[rune]rune{
		'B': '8',
		'E': '3',
		'G': '6',
		'H': '4',
		'I': '1',
		'L': '7',
		'O': '0',
		'S': '5',
		'Z': '2',
	}
	newLetter, exists := beghiloszRuneMap[letter]
	if !exists {
		return letter
	} else {
		return newLetter
	}
}

func Beghilosz(input string) string {
	lines := strings.Split(input, "\n")
	encodedLines := make([]string, len(lines))
	for lineIndex, line := range lines {
		words := strings.Fields(line)
		encodedWords := make([]string, len(words))
		for i, word := range words {
			word = strings.ToUpper(word)
			if !isSpecialWord(word) {
				word = strings.Map(beghiloszMapping, word)
				word = reverseString(word)
			}
			encodedWords[i] = word
		}
		encodedWords = reverseSlice(encodedWords)
		encodedLines[lineIndex] = strings.Join(encodedWords, " ")
	}
	return strings.Join(reverseSlice(encodedLines), "\n")
}

func isSpecialWord(word string) bool {
	isHashtag := strings.HasPrefix(word, "#")
	isUsername := strings.HasPrefix(word, "@")
	return isHashtag || isUsername
}

func reverseString(input string) string {
	var result string
	for _, letter := range input {
		result = string(letter) + result
	}
	return result
}

func reverseSlice(input []string) []string {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
	return input
}
