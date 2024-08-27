package palindrome_number_9

import "strconv"

func isPalindrome(x int) bool {

	// negative numbers cannot be palindromes
	if x < 0 {
		return false
	}

	// convert to string and compare first and last elements
	str := strconv.Itoa(x)
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}
