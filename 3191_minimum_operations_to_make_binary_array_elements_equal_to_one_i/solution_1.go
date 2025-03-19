package _191_minimum_operations_to_make_binary_array_elements_equal_to_one_i

func minOperations(nums []int) int {
	n := len(nums)
	ops := 0

	for i := 0; i <= n-3; i++ {
		if nums[i] == 0 {
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			ops++
		}
	}

	if nums[n-1] == 0 || nums[n-2] == 0 {
		return -1
	}

	return ops
}
