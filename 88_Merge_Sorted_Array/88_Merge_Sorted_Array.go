package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == m {
		return
	}

	if len(nums1) == n {
		copy(nums1, nums2)
	}

	i, j, k := m-1, n-1, m+n-1

	for j >= 0 && i >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

	fmt.Println(nums1)
}

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	nums2 := []int{2, 5, 6}
//
// 	merge(nums1, 3, nums2, 3)
// }
