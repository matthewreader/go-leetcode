package remove_element_27

func removeElement(nums []int, val int) int {
	b := nums[:0]
	for _, num := range nums {
		if num != val {
			b = append(b, num)
		}
	}
	return len(b)
}
