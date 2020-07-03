package utl

import (
	"strings"
	"unicode"
)

func Squish(input string) string {
	var pred rune

	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(pred) && unicode.IsSpace(r) {
			return -1
		}

		pred = r
		return r
	}, input)
}
