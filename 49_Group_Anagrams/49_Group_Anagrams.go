package main

import (
	"sort"
)

func sortString(s string) string {
	// Convert to slice of runes (to handle Unicode properly)
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, word := range strs {
		new := sortString(word)
		anagrams[new] = append(anagrams[new], word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, arrays := range anagrams {
		result = append(result, arrays)
	}
	return result
}

// func main() {
// 	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	fmt.Println(groupAnagrams(strs))
// }
