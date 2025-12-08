package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
	"unicode"
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

func readNextNum(reader *bufio.Reader) int {
	numStr := "";
	firstRune := rune(0);

	for {
		someRune, _, _ := reader.ReadRune();
		firstRune = someRune;
		if !unicode.IsDigit(someRune) {
			continue;
		} else {
			break;
		}
	}

	
	numStr += string(firstRune);

	for {
		someRune, _, _ := reader.ReadRune();
		if unicode.IsDigit(someRune) {
			numStr += string(someRune);
		} else {
			reader.UnreadRune();
			break;
		}
	}

	num, err := strconv.Atoi(numStr);
	if err != nil {
		panic(err);
	}

	return num;
}

func readOperations(reader *bufio.Reader) []string {

	output := []string{}
	for {
		someRune, _, err := reader.ReadRune();
		if (someRune == '+') || (someRune == '*') {
			output = append(output, string(someRune))
		} else if err == io.EOF {
			return output;
		} else {
			continue;
		}
	}
}

func finishMathProblems(problemMatrix [][]int, problemOperations []string) []int {

	N := len(problemMatrix);
	problems := len(problemMatrix[0]);

	output := []int{}
	for i := 0; i < problems; i++ {
		if problemOperations[i] == "*" {
			output = append(output, 1);
		} else {
			output = append(output, 0);
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < problems; j++ {
			if problemOperations[j] == "*" {
				output[j] *= problemMatrix[i][j];
			} else {
				output[j] += problemMatrix[i][j];
			}
		}
	}

	return output;
}

func main() {

	fileStr, _ := readFile("question6.input");
	reader := bufio.NewReader(strings.NewReader(fileStr));

	problemInputMatrix := [][]int{};
	problemOperations := []string{};
	newProblemLine := []int{};

	for {
		someRune, _, err := reader.ReadRune();
		if unicode.IsDigit(someRune) {
			reader.UnreadRune();
			num := readNextNum(reader);
			newProblemLine = append(newProblemLine, num);
		} else if someRune == '\n' {
			problemInputMatrix = append(problemInputMatrix, newProblemLine);
			newProblemLine = []int{};
		} else if (someRune == '+') || (someRune == '*') {
			problemOperations = append(problemOperations, string(someRune))
			problemOperations = append(problemOperations, readOperations(reader)...);
		} else if err == io.EOF {
			break;
		}
	}

	solutions := finishMathProblems(problemInputMatrix, problemOperations);
	total := 0
	for i := 0; i < len(solutions); i++ {
		total += solutions[i];
	}

	fmt.Println(total);
}
