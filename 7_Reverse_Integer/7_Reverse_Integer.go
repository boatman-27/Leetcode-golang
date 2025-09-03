package main

import "fmt"

const (
	INT_MAX = 1<<31 - 1 //  2147483647
	INT_MIN = -1 << 31  // -2147483648
)

func reverse(x int) int {
	var res int
	for x != 0 {
		lastDig := x % 10
		x /= 10
		res = res*10 + lastDig
	}

	if res > INT_MAX || res < INT_MIN {
		return 0
	}
	return res
}

func main() {
	fmt.Println(reverse(1534236469))
}
