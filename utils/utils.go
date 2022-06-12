// Package utils contains helper functions that are not directly related to any
// encoding.
package utils

import (
	"strings"
)

// IsSpecialWord return true if word is a hashtag or a username.
func IsSpecialWord(word string) bool {
	isHashtag := strings.HasPrefix(word, "#")
	isUsername := strings.HasPrefix(word, "@")
	return isHashtag || isUsername
}

// ReverseString reverse rune order in str.
func ReverseString(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
