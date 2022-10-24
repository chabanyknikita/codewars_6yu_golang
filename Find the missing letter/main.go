package main

func MissingLetter(chars []rune) rune {
	var last int

	for _, r := range chars {
		if last != 0 && int(r)-last > 1 {
			return rune(r - 1)
		}
		last = int(r)
	}

	return rune(last)
}
