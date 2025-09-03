package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}

	for _, char := range t {
		if count[char] == 0 {
			return false
		}
		count[char]--
	}

	return true
}

// func main() {
// 	s := "anagram"
// 	t := "nagaram"
//
// 	fmt.Println(isAnagram(s, t))
// }
