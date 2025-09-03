package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		midPoint := (left + right) / 2
		if nums[midPoint] == target {
			return midPoint
		}

		if target < nums[midPoint] {
			right = midPoint - 1
		} else {
			left = midPoint + 1
		}
	}

	return -1
}

// func main() {
// 	nums := []int{-1, 0, 3, 5, 9, 12}
// 	target := 9
//
// 	fmt.Println(search(nums, target))
// }
