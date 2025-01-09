package best_sightseeing_pair_1014

func maxScoreSightseeingPair(values []int) int {
	score := 0
	maxI := values[0]

	for j := 1; j < len(values); j++ {
		score = max(score, maxI+values[j]-j)
		maxI = max(maxI, values[j]+j)
	}

	return score
}
