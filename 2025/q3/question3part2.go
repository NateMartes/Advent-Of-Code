package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func readFile(name string) ([]string, error) {

	content, err := os.Open(name)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", name, err)
		os.Exit(1)
	}
	defer content.Close()
	var lines[]string
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getMaxJoltage(bank string) (int, error) {

	startIndex := 0
	padding := 11
	result := ""
	for padding >= 0 {
		maxVal := 0
		nextStart := startIndex
		limit := len(bank) - padding
		if (limit < 0) {
			break
		}
		for i := startIndex; i < limit; i++ {
			val, _ := strconv.Atoi(string(bank[i]))
			if (val <= maxVal) {
				continue
			}
			maxVal = max(val, maxVal)
			nextStart = i
		}
		result += strconv.Itoa(maxVal)
		padding--
		startIndex = nextStart + 1
	}

	return strconv.Atoi(result)

}

func main() {
	file := "question3.input"
	banks, _ := readFile(file)
	result := 0
	for i := 0; i < len(banks); i++ {
		max, _ := getMaxJoltage(banks[i])
		fmt.Println(max)
		result += max
	}
	fmt.Println(result)
}

