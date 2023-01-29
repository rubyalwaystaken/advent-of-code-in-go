package day_1

import (
	"log"
	"sort"
	"strconv"
)

func GetTotalCaloriesOfTopElvesSummed(calorieList []string, numberOfElvesToGet int) (summedCalories float64) {
	caloriesOfTopElves := GetTotalCaloriesOfTopElves(calorieList, numberOfElvesToGet)

	for _, calories := range caloriesOfTopElves {
		summedCalories += calories
	}

	return
}

func GetTotalCaloriesOfTopElves(calorieList []string, numberOfElvesToGet int) []float64 {
	var caloriesOfCurrentElf float64
	var totalCalorieList []float64

	for _, line := range calorieList {
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

func convertLineToFloat(line string) (float64, error) {
	return strconv.ParseFloat(line, 64)
}
