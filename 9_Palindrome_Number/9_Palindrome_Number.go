package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	org := x
	var rev int
	for x != 0 {
		last := x % 10
		x /= 10
		rev = rev*10 + last
	}

	return rev == org
}
