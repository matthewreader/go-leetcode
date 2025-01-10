package _8_find_the_index_of_the_first_occurrence_in_a_string

// Sliding window approach

func strStr(haystack string, needle string) int {
	needleLength := len(needle)
	haystackLength := len(haystack)

	if needleLength == 0 {
		return 0
	}

	for i := 0; i <= haystackLength-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}
	return -1
}
