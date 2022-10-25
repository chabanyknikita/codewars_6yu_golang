package main

func Persistence(n int) int {
	if n < 10 {
		return 0
	}
	mul := 1
	for n > 0 {
		mul *= n % 10
		n /= 10
	}
	return 1 + Persistence(mul)
}
