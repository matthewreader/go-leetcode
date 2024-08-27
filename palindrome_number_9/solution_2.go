package palindrome_number_9

func isPalindrome_2(x int) bool {
	// negative numbers cannot be palindromes
	if x < 0 {
		return false
	}

	// reverse the number
	original := x
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}

	return original == reversed
}
