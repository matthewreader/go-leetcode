package find_the_winner_of_the_circular_game_1823

func findTheWinner(n int, k int) int {
	// Create a list of n people
	people := make([]int, n)
	for i := 0; i < n; i++ {
		people[i] = i + 1
	}

	// Start the game
	// The game starts with the first person
	currentPerson := 0
	for len(people) > 1 {
		// Find the kth person
		kthPerson := (currentPerson + k - 1) % len(people)
		people = append(people[:kthPerson], people[kthPerson+1:]...)
		currentPerson = kthPerson
	}

	return people[0]
}
