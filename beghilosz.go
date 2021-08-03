package encodor

import (
	"strings"
)

func beghilosz_mapping(letter rune) rune {
	var char_map = map[rune]rune{
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
	new_letter, exists := char_map[letter]
	if !exists {
		return letter
	} else {
		return new_letter
	}
}

func Beghilosz(input string) string {
	lines := strings.Split(input, "\n")
	encoded_lines := make([]string, len(lines))
	for lineIndex, line := range lines {
		words := strings.Fields(line)
		encoded_words := make([]string, len(words))
		for i, word := range words {
			word = strings.ToUpper(word)
			if !isSpecialWord(word) {
				word = strings.Map(beghilosz_mapping, word)
				word = reverseString(word)
			}
			encoded_words[i] = word
		}
		encoded_words = reverseSlice(encoded_words)
		encoded_lines[lineIndex] = strings.Join(encoded_words, " ")
	}
	return strings.Join(reverseSlice(encoded_lines), "\n")
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
