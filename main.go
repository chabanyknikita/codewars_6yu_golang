import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	res := ""
	for _, word := range strings.Split(text, " ") {
		for i, letter := range word {
			switch i {
			case 0:
				res += string(strconv.Itoa(int(letter)))
			case 1:
				res += string(word[len(word)-1])
			case len(word) - 1:
				res += string(word[1])
			default:
				res += string(letter)
			}
		}
		res += " "
	}
	return strings.TrimSpace(res)
}
