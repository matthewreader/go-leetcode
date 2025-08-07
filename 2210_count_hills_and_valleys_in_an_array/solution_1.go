package _210_count_hills_and_valleys_in_an_array

func countHillValley(nums []int) int {
	count := 0
	n := len(nums)

	for i := 1; i < n-1; i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		prev := i - 1
		for prev >= 0 && nums[prev] == nums[i] {
			prev--
		}

		next := i + 1
		for next < n && nums[next] == nums[i] {
			next++
		}

		if prev >= 0 && next < n {
			if (nums[i] > nums[prev] && nums[i] > nums[next]) || (nums[i] < nums[prev] && nums[i] < nums[next]) {
				count++
			}
		}
	}
	return count
}
