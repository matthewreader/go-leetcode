package water_bottles_1518

func numWaterBottles(numBottles int, numExchange int) int {
	// Drink all water bottles at first step, then start exchanging
	fullBottles := 0
	emptyBottles := numBottles
	bottlesDrank := numBottles

	for emptyBottles >= numExchange {
		// Exchange the bottles
		fullBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange

		// Drink the full bottles
		bottlesDrank += fullBottles
		emptyBottles += fullBottles
		fullBottles = 0
	}
	return bottlesDrank
}
