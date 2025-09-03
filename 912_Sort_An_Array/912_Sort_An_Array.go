package main

func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	midpoint := len(nums) / 2
	left := nums[:midpoint]
	right := nums[midpoint:]

	return merge(sortArray(left), sortArray(right))
}

func merge(left, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]

	}

	return result
}
