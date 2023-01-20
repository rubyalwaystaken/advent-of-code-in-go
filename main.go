package main

import (
	day_1 "adventofcode/challenge_day_1"
	filereader "adventofcode/myutilities"
	"fmt"
)

func main() {
	input := filereader.ReadFile("files/challenge_1.txt")
	output := day_1.GetTotalCaloriesOfTopElvesSummed(input, 3)
	fmt.Println(output)
}
