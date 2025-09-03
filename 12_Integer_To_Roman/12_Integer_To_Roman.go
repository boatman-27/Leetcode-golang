package main

import "fmt"

func intToRoman(num int) string {
	var res string

	romanValue := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for _, symbol := range romans {
		value := romanValue[symbol]
		count := 0
		for num >= value {
			num -= value
			count++
		}

		for count > 0 {
			res += symbol
			count--
		}
	}

	return res
}

func main() {
	num := 3749
	fmt.Println(intToRoman(num))
}
