package main

import (
	"fgo24-go-append-slice/utils"
	"fmt"
)

func main() {
	scores1 := []int{50, 75, 66, 20, 32, 90}
	scores2 := []int{50, 75, 20, 32, 90, 66}
	scores3 := []int{50, 66, 75, 20, 32, 90}
	fmt.Println(utils.AppendSlice(scores1))
	fmt.Println(utils.AppendSlice(scores2))
	fmt.Println(utils.AppendSlice(scores3))
}
