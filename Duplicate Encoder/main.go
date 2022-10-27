package main

import (
	"strings"
	"unicode"
)

func DuplicateEncode(word string) string {
	encoded := strings.Builder{}
	occurrences := make(map[rune]int)
	for _, character := range word {
		occurrences[unicode.ToLower(character)]++
	}

	for _, character := range word {
		if occurrences[unicode.ToLower(character)] > 1 {
			encoded.WriteRune(')')
		} else {
			encoded.WriteRune('(')
		}
	}
	return encoded.String()
}
