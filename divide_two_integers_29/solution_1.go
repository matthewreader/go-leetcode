package divide_two_integers_29

/* Note, cannot use multiply, divide, or mod operator. */
func divide(dividend int, divisor int) int {
	res := 0
	absDividend := abs(dividend)
	absDivisor := abs(divisor)

	if divisor == 0 {
		return 0
	}

	/* This seems to be a strange bug in Leetcode? */
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	if divisor == -1 {
		return -dividend
	}

	if divisor == 1 {
		return dividend
	}

	for absDivisor <= absDividend {
		absDividend -= absDivisor
		res++
	}
	if dividend < 0 && divisor < 0 || dividend > 0 && divisor > 0 {
		return res
	}
	return -res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
