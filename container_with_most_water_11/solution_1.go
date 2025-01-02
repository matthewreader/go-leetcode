package container_with_most_water_11

// Note: the following code, while correct, is not the optimal solution.
// This solution times out the auto grader on LeetCode.

func maxArea(height []int) int {
	m := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			m = max(m, (j-i)*min(height[j], height[i]))
		}
	}
	return m
}
