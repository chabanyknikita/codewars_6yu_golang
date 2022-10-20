package main

import (
	"fmt"
	"strings"
	"unicode"
)

func shiftLetters(s string, n int) string {
	var result []string
	for _, ch := range s {
		newLetter := ch
		if unicode.IsLetter(ch) {
			newLetter = int32(int(ch) + n)
			if newLetter > 'z' {
				newLetter = ch - (26 - int32(n))
			}
		}
		result = append(result, string(newLetter))
	}
	return strings.Join(result, "")
}

func replaceDigitsWithComplements(s string) string {
	var result []string
	for _, ch := range s {
		newDigit := ch
		if unicode.IsDigit(ch) {
			newDigit = (57 - ch) + 48
		}
		result = append(result, string(newDigit))
	}
	return strings.Join(result, "")
}

func alternateCase(str string) string {
	chars := []rune{}
	for index, r := range str {
		if index%2 == 0 {
			chars = append(chars, unicode.ToUpper(r))
		} else {
			chars = append(chars, unicode.ToLower(r))
		}
	}
	return string(chars)
}

func reverseString(str string) string {
	var result []string
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, string(str[i]))
	}
	return strings.Join(result, "")
}

func PlayPass(s string, n int) string {
	return reverseString(alternateCase(replaceDigitsWithComplements(shiftLetters(s, n))))
}
func main() {
	fmt.Println(PlayPass("4897 NkTrC Hq fT67 GjV Pq [P OqTh gOcE CoPcTi [O", 0))
}
