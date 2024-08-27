package intersection_of_two_arrays_349

func intersection(nums1 []int, nums2 []int) []int {
	// Similar to solution to 350, but use a bool to indicate if the number
	// has been accounted for.
	nums1Map := make(map[int]bool)
	for _, num := range nums1 {
		nums1Map[num] = true
	}

	var result []int
	for _, num := range nums2 {
		if nums1Map[num] {
			result = append(result, num)
			nums1Map[num] = false
		}
	}

	return result

}
