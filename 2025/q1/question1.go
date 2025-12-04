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

func main() {
	file := "question1.input"
	turns, _ := readFile(file)

	LEFT_IDENT := "L"
	NUM_UPPER_BOUND := 99
	NUM_LOWER_BOUND := 0

	curr := 50
	result := 0

	for i := 0; i < len(turns); i++ {
		turn := turns[i]

		direction := turn[0:1]

		// Assume spin is to the right
		operator := func(a int, b int) int { return a + b }
		if (direction == LEFT_IDENT) {
			operator = func(a int, b int) int { return a - b }
		}

		value, err := strconv.Atoi(turn[1:len(turn)])
		if err != nil {
			panic(err)
		}

		tmp := value
		for tmp >= 1 {
			curr = operator(curr, 1)
			if curr < NUM_LOWER_BOUND {
				curr = NUM_UPPER_BOUND
			}
			if curr > NUM_UPPER_BOUND {
				curr = NUM_LOWER_BOUND
			}
			if (curr == 0) {
				result += 1
			}
			tmp--
		}

	}

	fmt.Println(result)
}
