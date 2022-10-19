package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DigPow(n, p int) int {
	first_res := 0.0
	arr_string := strings.Split(strconv.Itoa(n), "")
	for i := 0; i < len(arr_string); i++ {
		num, _ := strconv.Atoi(arr_string[i])
		first_res += math.Pow(float64(num), float64(p+i))
	}
	x := math.Pow(float64(n), float64(p))
	if first_res == x {
		return p
	} else if int(first_res)%n == 0 {
		return int(first_res) / n
	} else {
		return -1
	}
}
func main() {
	fmt.Println(DigPow(46288, 3))
}
