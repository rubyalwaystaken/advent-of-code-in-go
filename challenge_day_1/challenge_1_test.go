package day_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GivenCalorieList_WhenGetTotalCaloriesOfTopElf_ThenReturnCorrectValue(t *testing.T) {
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

	expectedMaxCalories := []float64{24000.0}
	givenMaxCalories := GetTotalCaloriesOfTopElves(calorieList, 1)

	assert.Equal(t, expectedMaxCalories, givenMaxCalories)
}

func Test_GivenCalorieList_WhenGetTotalCaloriesOfTopThreeElves_ThenReturnCorrectValue(t *testing.T) {
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

10000

`

	expectedMaxCalories := []float64{10000.0, 11000.0, 24000.0}
	givenMaxCalories := GetTotalCaloriesOfTopElves(calorieList, 3)

	assert.Equal(t, expectedMaxCalories, givenMaxCalories)
}

func Test_GivenCalorieList_WhenGetTotalCaloriesOfTopThreeElvesSummed_ThenReturnCorrectValue(t *testing.T) {
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

10000

`

	expectedSummedCalories := 45000.0
	givenSummedCalories := GetTotalCaloriesOfTopElvesSummed(calorieList, 3)

	assert.Equal(t, expectedSummedCalories, givenSummedCalories)
}
