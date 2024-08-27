package minimum_difference_between_largest_and_smallest_value_in_three_moves_1509

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)

	minDiff := math.MaxInt32 // use math.MaxInt32 instead of int, as int gets initialized to 0.
	for i := 0; i < 4; i++ {
		minDiff = min(minDiff, nums[len(nums)-4+i]-nums[i])
	}

	return minDiff
}
