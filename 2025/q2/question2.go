package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func readFile(name string) (string, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	str := string(b)
    return str, nil
}

func getRanges(rangelist string) ([]string) {
    return strings.Split(rangelist, ",")
}

func main() {
	file := "question2.input"
	fileString, _ := readFile(file)
	ranges := getRanges(strings.Trim(fileString, "\n"))
	invalids := 0
	for i := 0; i < len(ranges); i++ {
		currRange := ranges[i]
		currRangeSplit := strings.Split(currRange, "-")
		lowerRange, err := strconv.Atoi(currRangeSplit[0])
		if err != nil {
			panic(err)
		}
		upperRange, err := strconv.Atoi(currRangeSplit[1])
		if err != nil {
			panic(err)
		}
		curr := lowerRange
		for curr <= upperRange {
			currString := strconv.Itoa(curr)
			if len(currString) % 2 == 0 {
				half := len(currString) / 2
				substring1, _ := strconv.Atoi(currString[:half])
				substring2, _ := strconv.Atoi(currString[half:])
				if substring1 == substring2 {
					invalids += curr
				}
			}
			curr++
		}
	}
	fmt.Println(invalids)
}

