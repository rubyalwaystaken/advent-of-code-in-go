package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func getTotalCaloriesOfElfCarryingTheMost(calorieList string) (totalCalories float64) {
	lineArray, error := convertToStringArray(calorieList)

	if error != nil {
		log.Fatal(error)
	}

	caloriesOfCurrentElf := 0.0

	for _, line := range lineArray {
		if line == "\n" || line == "" {
			if caloriesOfCurrentElf > totalCalories {
				totalCalories = caloriesOfCurrentElf
			}
			caloriesOfCurrentElf = 0.0
			continue
		}

		convertedLine, error := convertLineToFloat(line)

		if error != nil {
			log.Fatal(error)
		}

		caloriesOfCurrentElf += convertedLine
	}

	return totalCalories
}

func convertToStringArray(calorieList string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(calorieList))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	return
}

func convertLineToFloat(line string) (float64, error) {
	return strconv.ParseFloat(line, 64)
}

func main() {
}
