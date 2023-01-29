package day_2

import (
	filereader "adventofcode/myutilities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = filereader.ReadFile("../files/test_challenge_2.txt")

func Test_GivenStrategyGuide_WhenCalculateTotalScoreByFirstStrategy_ThenReturnCorrectResult(t *testing.T) {
	expected := 15
	actual := CalculateTotalScoreByFirstStrategy(input)

	assert.Equal(t, expected, actual)
}

func Test_GivenStrategyGuide_WhenCalculateTotalScoreBySecondStrategy_ThenReturnCorrectResult(t *testing.T) {
	expected := 12
	actual := CalculateTotalScoreBySecondStrategy(input)

	assert.Equal(t, expected, actual)
}
