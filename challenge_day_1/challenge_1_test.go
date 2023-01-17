package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenCalorieList_WhenGetTotalCaloriesOfElfCarryingTheMost_ThenReturnCorrectValue(t *testing.T) {
	calorieList :=
		`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	expectedMaxCalories := 24000.0
	givenMaxCalories := getTotalCaloriesOfElfCarryingTheMost(calorieList)

	assert.Equal(t, expectedMaxCalories, givenMaxCalories)
}
