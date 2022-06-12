// Package utils contains helper functions that are not directly related to any
// encoding.
package utils

import (
	"bytes"
)

// IsSpecialWord return true if word is a hashtag or a username.
func IsSpecialWord(word []byte) bool {
	isHashtag := bytes.HasPrefix(word, []byte("#"))
	isUsername := bytes.HasPrefix(word, []byte("@"))
	return isHashtag || isUsername
}

// ReverseLetters reverse runes order in buf.
func ReverseLetters(buf []byte) []byte {
	rns := bytes.Runes(buf)
	for i := 0; i < len(rns)/2; i++ {
		rns[i], rns[len(rns)-1-i] = rns[len(rns)-1-i], rns[i]
	}
	return []byte(string(rns))
}

// ReverseLines reverse lines order in slice of lines.
func ReverseLines(slice [][]byte) [][]byte {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
}
