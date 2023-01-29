package day_1

import (
	filereader "adventofcode/myutilities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCalorieList = filereader.ReadFile("../files/test_challenge_1.txt")

func Test_GivenCalorieList_WhenGetTotalCaloriesOfTopElf_ThenReturnCorrectValue(t *testing.T) {
	expectedMaxCalories := []float64{24000.0}
	givenMaxCalories := GetTotalCaloriesOfTopElves(testCalorieList, 1)

	assert.Equal(t, expectedMaxCalories, givenMaxCalories)
}

func Test_GivenCalorieList_WhenGetTotalCaloriesOfTopThreeElves_ThenReturnCorrectValue(t *testing.T) {
	expectedMaxCalories := []float64{10000.0, 11000.0, 24000.0}
	givenMaxCalories := GetTotalCaloriesOfTopElves(testCalorieList, 3)

	assert.Equal(t, expectedMaxCalories, givenMaxCalories)
}

func Test_GivenCalorieList_WhenGetTotalCaloriesOfTopThreeElvesSummed_ThenReturnCorrectValue(t *testing.T) {
	expectedSummedCalories := 45000.0
	givenSummedCalories := GetTotalCaloriesOfTopElvesSummed(testCalorieList, 3)

	assert.Equal(t, expectedSummedCalories, givenSummedCalories)
}
