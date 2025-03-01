package _460_apply_operations_to_an_array

func applyOperations(nums []int) []int {
	numsLen := len(nums)

	for i := 0; i < numsLen-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	nums = moveZeroes(nums)
	return nums
}

func moveZeroes(nums []int) []int {
	lastNonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroIndex] = nums[i]
			lastNonZeroIndex++
		}
	}

	// Set remaining element to zero.
	for i := lastNonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}
