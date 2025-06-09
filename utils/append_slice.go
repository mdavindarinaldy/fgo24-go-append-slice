package utils

func AppendSlice(scores []int) []int {
	newScores := []int{}
	for i := range scores {
		newScores = append(newScores, scores[i])
		if scores[i] == 66 {
			newScores = append(newScores, 88)
		}
	}
	return newScores
}
