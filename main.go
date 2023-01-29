package main

import (
	day_2 "adventofcode/challenge_day_2"
	filereader "adventofcode/myutilities"
	"fmt"
)

func main() {
	input := filereader.ReadFile("files/challenge_2.txt")
	output := day_2.CalculateTotalScoreByFirstStrategy(input)
	fmt.Println(output)
	output = day_2.CalculateTotalScoreBySecondStrategy(input)
	fmt.Println(output)
}
