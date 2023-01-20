package day_1

import (
	"bufio"
	"log"
	"sort"
	"strconv"
	"strings"
)

func GetTotalCaloriesOfTopElvesSummed(calorieList string, numberOfElvesToGet int) (summedCalories float64) {
	caloriesOfTopElves := GetTotalCaloriesOfTopElves(calorieList, numberOfElvesToGet)

	for _, calories := range caloriesOfTopElves {
		summedCalories += calories
	}

	return
}

func GetTotalCaloriesOfTopElves(calorieList string, numberOfElvesToGet int) []float64 {
	lineArray, error := convertToStringArray(calorieList)

	if error != nil {
		log.Fatal(error)
	}

	var caloriesOfCurrentElf float64
	var totalCalorieList []float64

	for _, line := range lineArray {
		if line == "\n" || line == "" {
			totalCalorieList = append(totalCalorieList, caloriesOfCurrentElf)
			caloriesOfCurrentElf = 0.0
			continue
		}

		convertedLine, error := convertLineToFloat(line)

		if error != nil {
			log.Fatal(error)
		}

		caloriesOfCurrentElf += convertedLine
	}

	sort.Float64s(totalCalorieList)

	return totalCalorieList[len(totalCalorieList)-numberOfElvesToGet:]
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
