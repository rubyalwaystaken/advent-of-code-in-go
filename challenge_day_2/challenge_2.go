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

var winMap = map[string]string{
	"A": "Y", // Paper beats Rock
	"B": "Z", // Scissors beat Paper
	"C": "X", // Rock beats Scissors
}

var loseMap = map[string]string{
	"A": "Z", // Scissors loses to Rock
	"B": "X", // Rock loses to Paper
	"C": "Y", // Paper loses to Scissors
}

func CalculateTotalScoreByFirstStrategy(input []string) (totalScore int) {
	for _, line := range input {
		columns := splitByColumn(line)

		var roundScore = 0

		if drawMap[columns[0]] == columns[1] {
			roundScore = 3
		} else if winMap[columns[0]] == columns[1] {
			roundScore = 6
		}

		totalScore += roundScore
		totalScore += shapeScore[columns[1]]
	}

	return
}

func CalculateTotalScoreBySecondStrategy(input []string) (totalScore int) {
	for _, line := range input {
		columns := splitByColumn(line)

		var roundScore = 0
		var shapeToChoose = ""

		switch moveToMake := columns[1]; moveToMake {
		case "X": // Lose
			shapeToChoose = loseMap[columns[0]]
		case "Y": // Draw
			shapeToChoose = drawMap[columns[0]]
			roundScore = 3
		case "Z": // Win
			shapeToChoose = winMap[columns[0]]
			roundScore = 6
		}

		totalScore += roundScore
		totalScore += shapeScore[shapeToChoose]
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
