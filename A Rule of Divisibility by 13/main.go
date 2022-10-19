package main

func Thirt(n int) int {
	factor := []int{1, 10, 9, 12, 3, 4}
	sum := n
	before := 0
	for sum != before {
		before = sum
		sum = 0
		i := 0
		n = before
		for n > 0 {
			sum += (n % 10) * factor[i%6]
			i++
			n = n / 10
		}
	}
	return sum
}
