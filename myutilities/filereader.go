package filereader

import (
	"io/ioutil"
	"log"
)

func ReadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
