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

func readOperations(reader *bufio.Reader) []rune {

	output := []rune{}
	for {
		someRune, _, err := reader.ReadRune();
		if (someRune == '+') || (someRune == '*') {
			output = append(output, someRune)
		} else if err == io.EOF {
			return output;
		} else {
			continue;
		}
	}
}


func getCephalopodNums(problemInputMatrix [][]rune) [][]int {

	output := [][]int{};

	cols := len(problemInputMatrix[0]);
	rows := len(problemInputMatrix);

	currCol := []int{}

	for i := 0; i < cols; i++ {

		num := "";
		for j := 0; j < rows; j++ {
			curr := problemInputMatrix[j][i];
			if unicode.IsDigit(curr) {
				num += string(curr)
			}
		}

		if num == "" {
			output = append(output, currCol);
			currCol = []int{};
		} else {
			numInt, _ := strconv.Atoi(num);
			currCol = append(currCol, numInt);
		}

	}

	return append(output, currCol);
}

func finishMathProblems(problemMatrix [][]int, problemOperations []rune) []int {

	problems := len(problemOperations);
	output := []int{};

	for i := 0; i < problems; i++ {
		nums := problemMatrix[i];
		operation := problemOperations[i]
		
		total := 1;
		if operation == '+' {
			total = 0;
		}

		for j := 0; j < len(nums); j++ {
			if operation == '*' {
				total *= nums[j];
			} else {
				total += nums[j];
			}
		}

		output = append(output, total);
	}

	return output;
}

func main() {

	fileStr, _ := readFile("question6.input");
	reader := bufio.NewReader(strings.NewReader(fileStr));


	problemInputMatrix := [][]rune{};
	problemOperations := []rune{};
	newProblemLine := []rune{};

	for {

		someRune, _, err := reader.ReadRune();
		if unicode.IsDigit(someRune) || someRune == ' ' {

			newProblemLine = append(newProblemLine, someRune);

		} else if someRune == '\n' {

			problemInputMatrix = append(problemInputMatrix, newProblemLine);
			newProblemLine = []rune{};

		} else if (someRune == '+') || (someRune == '*') {

			problemOperations = append(problemOperations, someRune)
			problemOperations = append(problemOperations, readOperations(reader)...);

		} else if err == io.EOF {
			break;
		}
	}

	fixedMatrix := getCephalopodNums(problemInputMatrix);

	solutions := finishMathProblems(fixedMatrix, problemOperations);
	total := 0
	for i := 0; i < len(solutions); i++ {
		total += solutions[i];
	}

	fmt.Println(total);
}
