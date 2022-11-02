package main

import (
	"strings"
)

func Encode(s string, n int) string {
	rails := make([]string, n)
	i, sign := 0, 1
	for _, c := range s {
		rails[i] += string(c)
		if val := i + sign; val == -1 || val == n {
			sign *= -1
		}
		i += sign
	}
	return strings.Join(rails, "")
}

func Decode(s string, n int) string {
	postDict := 2*n - 2
	k, l := 0, len(s)
	decode := make([]byte, l)
	for i := 0; i < n; i++ {
		offset := postDict - 2*i
		for j := i; j < l; j += postDict {
			if k < l {
				decode[j] = s[k]
				k++
			}
			if i != 0 && i != n-1 && j+offset < l {
				decode[j+offset] = s[k]
				k++
			}
		}
	}
	return string(decode)

}
