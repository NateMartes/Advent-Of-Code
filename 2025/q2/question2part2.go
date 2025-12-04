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
			currStringLength := len(currString)

			for j := 1; j < currStringLength; j++ {

				currSubStringTest := currString[0:j]
				currSubStringTestLength := len(currSubStringTest)
				modVal := currStringLength % currSubStringTestLength
				if modVal == 0 || modVal == currStringLength {
					isRep := true
					for k := j; k < currStringLength; k += currSubStringTestLength {
						if currSubStringTest != currString[k:k + currSubStringTestLength] {
							isRep = false
							break
						}
					}

					if isRep {
						fmt.Println(curr)
						invalids += curr
						break
					}
				}
			}

			curr++
		}
	}

	fmt.Println(invalids)
}


