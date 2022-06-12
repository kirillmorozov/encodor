// Package zalgo provides Encode function that turns your text into cryptic and
// glitch zalgo text.
//
// Zalgo text is digital text that has been modified with combining characters,
// Unicode symbols used to add diacritics above or below letters, to appear
// frightening or glitchy.
package zalgo

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"unicode"

	"github.com/kirillmorozov/encodor/utils"
)

const (
	minDiacritics = 1
	maxDiacritics = 5
)

var highDiacritics = []rune{
	'\u030d', '\u030e', '\u0304', '\u0305', '\u033f', '\u0311', '\u0306',
	'\u0310', '\u0352', '\u0357', '\u0351', '\u0307', '\u0308', '\u030a',
	'\u0342', '\u0343', '\u0344', '\u034a', '\u034b', '\u034c', '\u0303',
	'\u0302', '\u030c', '\u0350', '\u0300', '\u0301', '\u030b', '\u030f',
	'\u0312', '\u0313', '\u0314', '\u033d', '\u0309', '\u0363', '\u0364',
	'\u0365', '\u0366', '\u0367', '\u0368', '\u0369', '\u036a', '\u036b',
	'\u036c', '\u036d', '\u036e', '\u036f', '\u033e', '\u035b', '\u0346',
	'\u031a',
}

var midDiacritics = []rune{
	'\u0315', '\u031b', '\u0340', '\u0341', '\u0358', '\u0321', '\u0322',
	'\u0327', '\u0328', '\u0334', '\u0335', '\u0336', '\u034f', '\u035c',
	'\u035d', '\u035e', '\u035f', '\u0360', '\u0362', '\u0338', '\u0337',
	'\u0361', '\u0489',
}

var lowDiacritics = []rune{
	'\u0316', '\u0317', '\u0318', '\u0319', '\u031c', '\u031d', '\u031e',
	'\u031f', '\u0320', '\u0324', '\u0325', '\u0326', '\u0329', '\u032a',
	'\u032b', '\u032c', '\u032d', '\u032e', '\u032f', '\u0330', '\u0331',
	'\u0332', '\u0333', '\u0339', '\u033a', '\u033b', '\u033c', '\u0345',
	'\u0347', '\u0348', '\u0349', '\u034d', '\u034e', '\u0353', '\u0354',
	'\u0355', '\u0356', '\u0359', '\u035a', '\u0323',
}

// Encode will run the zalgo text generator against the input.
//
// Hashtags(words beginning with `#`) and mentions(words beginning with `@`) are
// left as is.
func Encode(reader io.Reader, writer io.Writer, diacritics int8) error {
	if (diacritics < minDiacritics) || (diacritics > maxDiacritics) {
		return errors.New("Incorrect number of diacritics, should be 1 <= diacritics <= 5")
	}
	input := bufio.NewScanner(reader)
	input.Split(bufio.ScanWords)
	output := bufio.NewWriter(writer)
	defer output.Flush()
	for input.Scan() {
		if scanErr := input.Err(); scanErr != nil {
			return scanErr
		}
		word := input.Bytes()
		if utils.IsSpecialWord(word) {
			_, writeErr := output.Write(input.Bytes())
			if writeErr != nil {
				return writeErr
			}
			continue
		}
		for _, r := range string(word) {
			_, writeErr := output.WriteRune(r)
			if writeErr != nil {
				return writeErr
			}
			if unicode.IsLetter(rune(r)) || unicode.IsDigit(rune(r)) {
				for i := int8(0); i < diacritics; i++ {
					_, writeErr := output.WriteRune(randZalgo(highDiacritics))
					if writeErr != nil {
						return writeErr
					}
					_, writeErr = output.WriteRune(randZalgo(midDiacritics))
					if writeErr != nil {
						return writeErr
					}
					_, writeErr = output.WriteRune(randZalgo(lowDiacritics))
					if writeErr != nil {
						return writeErr
					}
				}
			}
		}
	}
	return nil
}

// randZalgo gets a random char from a zalgo char table
func randZalgo(array []rune) rune {
	index := rand.Intn(len(array))
	return array[index]
}
