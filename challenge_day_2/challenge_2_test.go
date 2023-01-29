package day_2

import (
	filereader "adventofcode/myutilities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = filereader.ReadFile("../files/test_challenge_2.txt")

func Test_GivenStrategyGuide_WhenCalculateTotalScore_ThenReturnCorrectResult(t *testing.T) {
	expected := 15
	actual := CalculateTotalScore(input)

	assert.Equal(t, expected, actual)
}
