package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	arrayRunes := []rune(s)
	var result string
	var n int
	var backslash bool

	for i, char := range arrayRunes {
		if unicode.IsDigit(char) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) && unicode.IsDigit(arrayRunes[i-1]) && arrayRunes[i-2] != '\\' {
			return "", ErrInvalidString
		}
		if char == '\\' && !backslash {
			backslash = true
			continue
		}
		if backslash && unicode.IsLetter(char) {
			return "", ErrInvalidString
		}
		if backslash {
			result += string(char)
			backslash = false
			continue
		}
		if unicode.IsDigit(char) {
			n = int(char - '0')
			if n == 0 {
				result = result[:len(result)-1]
				continue
			}
			for j := 0; j < n-1; j++ {
				result += string(arrayRunes[i-1])
			}
			continue
		}
		result += string(char)
	}

	return result, nil
}
