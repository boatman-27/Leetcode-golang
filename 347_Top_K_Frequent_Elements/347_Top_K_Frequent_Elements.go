package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	}

	var result []int
	keys := make([]int, 0, len(freqs))

	for key := range freqs {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return freqs[keys[i]] > freqs[keys[j]]
	})

	for _, value := range keys {
		if len(result) == 2 {
			return result
		}
		result = append(result, value)

	}

	return result
}

// func main() {
// 	nums := []int{1, 1, 1, 2, 2, 3}
// 	k := 2
//
// 	fmt.Println(topKFrequent(nums, k))
// }
