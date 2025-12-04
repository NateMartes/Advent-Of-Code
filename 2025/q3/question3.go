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
	maxLeftBank := 0
	maxLeftBankIndex := 0
	maxRightBank := 0

	for i := 0; i < len(bank) - 1; i++ {
		val, _ := strconv.Atoi(string(bank[i]))
		if (val <= maxLeftBank) {
			continue
		}
		maxLeftBank = max(val, maxLeftBank)
		maxLeftBankIndex = i
	}

	for i := maxLeftBankIndex + 1; i < len(bank); i++ {
		val, _ := strconv.Atoi(string(bank[i]))
		maxRightBank = max(val, maxRightBank)
	}

	return strconv.Atoi(strconv.Itoa(maxLeftBank) + strconv.Itoa(maxRightBank))

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

