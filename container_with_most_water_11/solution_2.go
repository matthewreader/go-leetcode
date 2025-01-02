package container_with_most_water_11

// Note: Optimized code using two pointers.

func maxArea_2(height []int) int {
	m := 0
	left, right := 0, len(height)-1
	for left < right {
		width := right - left
		h := min(height[left], height[right])
		m = max(m, width*h)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return m
}
