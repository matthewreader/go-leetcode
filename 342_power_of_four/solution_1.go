package _42_power_of_four

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	// Check power of 2 and only set bit at even position
	return (n&(n-1)) == 0 && (n&0x55555555) != 0
}
