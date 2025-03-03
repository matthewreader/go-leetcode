package _161_partition_array_according_to_given_pivot

// Three pass solution, O(n) complexity
func pivotArray(nums []int, pivot int) []int {
	var lessThan []int
	var greaterThan []int
	var equalTo []int

	for _, num := range nums {
		if num < pivot {
			lessThan = append(lessThan, num)
		}
	}

	for _, num := range nums {
		if num == pivot {
			equalTo = append(equalTo, num)
		}
	}

	for _, num := range nums {
		if num > pivot {
			greaterThan = append(greaterThan, num)
		}
	}
	return append(append(lessThan, equalTo...), greaterThan...)
}
