package main

import "fmt"

func twoSum(nums []int, target int) []int {
	amountLeft := make(map[int]int)

	for index, value := range nums {
		remain := target - value
		if j, ok := amountLeft[remain]; ok {
			return []int{index, j}
		}
		amountLeft[value] = index
	}

	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}
