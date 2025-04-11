package _843_count_symmetric_integers

import "strconv"

func countSymmetricIntegers(low int, high int) int {
	integerCount := 0

	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		length := len(s)

		if length%2 != 0 {
			continue
		}

		half := length / 2
		firstHalfSum := 0
		secondHalfSum := 0

		for i := 0; i < half; i++ {
			firstHalfSum += int(s[i] - '0')
			secondHalfSum += int(s[i+half] - '0')
		}

		if firstHalfSum == secondHalfSum {
			integerCount++
		}
	}
	return integerCount
}
