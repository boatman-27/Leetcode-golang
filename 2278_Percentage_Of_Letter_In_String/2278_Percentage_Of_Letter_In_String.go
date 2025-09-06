package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			counter++
		}
	}
	return counter * 100 / len(s)
}

func main() {
	s := "foobar"
	var l byte = 157

	fmt.Println(percentageLetter(s, l))
}
