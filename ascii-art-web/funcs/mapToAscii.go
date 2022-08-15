package AsciiArt

import (
	"io/ioutil"
	"log"
	"strings"
)

func MapAscii(banner string) map[int]string {

	file, err := ioutil.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}

	indivLine := strings.Split(string(file), "\n")

	art := make(map[int]string)

	for i, j := 32, 0; i < 127; i++ {
		k := j + 9
		temp := ""

		for j < k {
			temp += indivLine[j+1] + "\n"
			j++
		}

		art[i] = temp
	}
	return art
}
