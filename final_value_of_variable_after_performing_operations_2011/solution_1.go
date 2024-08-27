package final_value_of_variable_after_performing_operations_2011

import (
	"regexp"
)

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, operation := range operations {
		if matched, _ := regexp.MatchString("\\+\\+", operation); matched {
			x++
		} else if matched, _ := regexp.MatchString("--", operation); matched {
			x--
		}
	}
	return x
}
