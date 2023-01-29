package filereader

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(fileName string) (lineArray []string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lineArray, err = convertToStringArray(string(content))
	if err != nil {
		log.Fatal(err)
	}

	return
}

func convertToStringArray(calorieList string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(calorieList))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	return
}
