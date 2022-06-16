package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sNew, substring string
	var num int
	var numIndicator, letterIndicator, shield bool
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		letter := s1[i]
		if string(letter) == "\\" && !shield {
			shield = true
			continue
		}
		switch {
		case shield:
			shield = false
			sNew += substring
			substring = string(letter)
			letterIndicator = true
		case unicode.IsDigit(letter):
			if !letterIndicator || numIndicator {
				return "", ErrInvalidString
			}
			num, _ = strconv.Atoi(string(letter))
			sNew += strings.Repeat(substring, num)
			numIndicator = true
			letterIndicator = false
		case letterIndicator:
			sNew += substring
			substring = string(letter)
			numIndicator = false
		default:
			substring = string(letter)
			letterIndicator = true
			numIndicator = false
		}
	}
	if letterIndicator {
		sNew += substring
	}
	return sNew, nil
}
