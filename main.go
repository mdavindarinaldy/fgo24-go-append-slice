package main

import (
	"fmt"
)

func main() {
	scores := []int{50, 75, 66, 20, 32, 90}
	var newScores [7]int
	for x := range scores {
		if scores[x] == 20 {
			newScores[x] = 88
			newScores[x+1] = scores[x]
		} else if x > 3 {
			newScores[x+1] = scores[x]
		} else {
			newScores[x] = scores[x]
		}
	}
	fmt.Println(newScores)
}
