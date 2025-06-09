package main

import (
	"fmt"
)

func main() {
	scores := []int{50, 75, 66, 20, 32, 90}
	var newScores [7]int
	for x := range scores {
		if scores[x] == 66 {
			newScores[x] = scores[x]
			newScores[x+1] = 88
		} else if x > 2 {
			newScores[x+1] = scores[x]
		} else {
			newScores[x] = scores[x]
		}
	}
	fmt.Println(newScores)
}
