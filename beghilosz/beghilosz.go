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
	"bufio"
	"bytes"
	"io"

	"github.com/kirillmorozov/encodor/utils"
)

// beghiloszMapper returns a digit that corresponds to a given rune in a
// calculator spelling or the rune itself.
func beghiloszMapper(r rune) rune {
	switch r {
	case 'B':
		return '8'
	case 'E':
		return '3'
	case 'G':
		return '6'
	case 'H':
		return '4'
	case 'I':
		return '1'
	case 'L':
		return '7'
	case 'O':
		return '0'
	case 'S':
		return '5'
	case 'Z':
		return '2'
	default:
		return r
	}
}

// Encode transforms text into calculator spelling.
//
// Hashtags(words beginning with `#`) and mentions(words beginning with `@`) are
// left as is.
func Encode(reader io.Reader, writer io.Writer) error {
	lineScanner := bufio.NewScanner(reader)
	lineScanner.Split(bufio.ScanLines)
	var line [][]byte
	var lines [][]byte
	// Encode
	for lineScanner.Scan() {
		if scanErr := lineScanner.Err(); scanErr != nil {
			return scanErr
		}
		line = nil
		wordScanner := bufio.NewScanner(bytes.NewReader(lineScanner.Bytes()))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			if scanErr := wordScanner.Err(); scanErr != nil {
				return scanErr
			}
			word := bytes.ToUpper(wordScanner.Bytes())
			if utils.IsSpecialWord(word) {
				line = append(line, utils.ReverseLetters(word))
			} else {
				line = append(line, bytes.Map(beghiloszMapper, word))
			}
		}
		lines = append(lines, utils.ReverseLetters(bytes.Join(line, []byte(" "))))
	}
	lines = utils.ReverseLines(lines)
	if _, writeErr := writer.Write(bytes.Join(lines, []byte("\n"))); writeErr != nil {
		return writeErr
	}
	return nil
}
