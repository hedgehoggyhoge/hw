package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	str := []rune(s)
	var stringNew string
	var number int
	var backSlash bool

	for i, item := range str {
		if unicode.IsDigit(item) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(item) && unicode.IsDigit(str[i-1]) && str[i-2] != '\\' {
			return "", ErrInvalidString
		}
		if item == '\\' && !backSlash {
			backSlash = true
			continue
		}
		if backSlash && unicode.IsLetter(item) {
			return "", ErrInvalidString
		}
		if backSlash {
			stringNew += string(item)
			backSlash = false
			continue
		}
		if unicode.IsDigit(item) {
			number = int(item - '0')
			if number == 0 {
				stringNew = stringNew[:len(stringNew)-1]
				continue
			}
			for j := 0; j < number-1; j++ {
				stringNew += string(str[i-1])
			}
			continue
		}
		stringNew += string(item)
	}
	return stringNew, nil
}
