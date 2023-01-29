package day_2

import (
	"log"
	"strings"
)

var shapeScore = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

var drawMap = map[string]string{
	"A": "X", // Rock
	"B": "Y", // Paper
	"C": "Z", // Scissors
}

var winCombo = map[string]string{
	"A": "Y", // Paper beats Rock
	"B": "Z", // Scissors beat Paper
	"C": "X", // Rock beats Scissors
}

func CalculateTotalScore(input []string) (totalScore int) {
	for _, line := range input {
		columns := splitByColumn(line)

		var roundScore = 0

		if drawMap[columns[0]] == columns[1] {
			roundScore = 3
		} else if winCombo[columns[0]] == columns[1] {
			roundScore = 6
		}

		totalScore += roundScore
		totalScore += shapeScore[columns[1]]
	}

	return
}

func splitByColumn(line string) (split []string) {
	split = strings.Fields(line)
	if len(split) != 2 {
		log.Fatalf("Line does not contain exactly 2 characters")
	}

	return
}
