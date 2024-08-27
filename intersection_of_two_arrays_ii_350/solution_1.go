package intersection_of_two_arrays_ii_350

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	for _, num := range nums1 {
		nums1Map[num]++
	}

	var result []int
	for _, num := range nums2 {
		if nums1Map[num] > 0 {
			result = append(result, num)
			nums1Map[num]--
		}
	}

	return result
}
