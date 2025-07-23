package _717_maximum_score_substring_removal

func maximumGain(s string, x int, y int) int {
	var first, second string
	var firstPoints, secondPoints int

	if x >= y {
		first, second = "ab", "ba"
		firstPoints, secondPoints = x, y
	} else {
		first, second = "ba", "ab"
		firstPoints, secondPoints = y, x
	}
	remaining, score1 := removeSubstring(s, first, firstPoints)
	_, score2 := removeSubstring(remaining, second, secondPoints)

	return score1 + score2
}

func removeSubstring(s string, target string, points int) (string, int) {
	stack := make([]byte, 0, len(s))
	score := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		stack = append(stack, char)

		// Check if we can form the target substring
		if len(stack) >= 2 &&
			stack[len(stack)-2] == target[0] &&
			stack[len(stack)-1] == target[1] {

			// Remove the two characters that form the target
			stack = stack[:len(stack)-2]
			score += points
		}
	}

	return string(stack), score
}
