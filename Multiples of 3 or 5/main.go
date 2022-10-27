package main

import "fmt"

func Multiple3And5(number int) int {
	var new_arr []int
	summary := 0
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			new_arr = append(new_arr, i)
		}
	}
	for _, i := range new_arr {
		summary += i
	}
	return summary
}
