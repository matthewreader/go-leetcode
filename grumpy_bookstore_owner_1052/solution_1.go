package grumpy_bookstore_owner_1052

func maxSatisfied(customers []int, grumpy []int, minutes int) int {

	// Calculate the number of satisfied customers before calculating the maximum number of additional satisfied
	// customers.
	satisfiedCustomers := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			satisfiedCustomers += customers[i]
		}
	}

	// Calculate the maximum number of additional satisfied customers given a window of n minutes.
	unsatisfiedCustomers := make([]int, len(customers))

	for i := 0; i < len(customers); i++ {
		unsatisfiedCustomers[i] = customers[i] * grumpy[i]
	}

	maxAdditionalSatisfiedCustomers := 0
	for i := 0; i < minutes; i++ {
		maxAdditionalSatisfiedCustomers += unsatisfiedCustomers[i]
	}

	currentAdditionalSatisfiedCustomers := maxAdditionalSatisfiedCustomers
	for i := minutes; i < len(customers); i++ {
		currentAdditionalSatisfiedCustomers += unsatisfiedCustomers[i] - unsatisfiedCustomers[i-minutes]
		if currentAdditionalSatisfiedCustomers > maxAdditionalSatisfiedCustomers {
			maxAdditionalSatisfiedCustomers = currentAdditionalSatisfiedCustomers
		}
	}

	return maxAdditionalSatisfiedCustomers + satisfiedCustomers
}
