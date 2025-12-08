package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
	"unicode"
	"strconv"
	"sort"
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

func getCephalopodNum(nums []string) []int {

}

func getCorrectNumsForProblems(problemInputMatrix [][]string) [][]int {
	output := [][]int{};
	cols := len(problemInputMatrix[0]);
	rows := len(problemInputMatrix);
	for i := 0; i < cols; i++ {
		nums := []string{}
		for j := 0; j < rows; j++ {
			nums = append(nums, problemInputMatrix[j][i]);
		}
		sort.Slice(nums, func(i, j int) bool {
			lenA := len(nums[i])
			lenB := len(nums[j])

			if lenA != lenB {
				return lenA < lenB
			}
			return nums[i] < nums[j]
		})
		fmt.Println(nums);
		res := getCephalopodNum(nums);
		fmt.Println(res);
	}

	return output;
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

	fileStr, _ := readFile("question6.testinput");
	reader := bufio.NewReader(strings.NewReader(fileStr));

	problemInputMatrix := [][]string{};
	problemOperations := []string{};
	newProblemLine := []string{};

	for {
		someRune, _, err := reader.ReadRune();
		if unicode.IsDigit(someRune) {
			reader.UnreadRune();
			num := readNextNum(reader);
			numStr := strconv.Itoa(num);
			newProblemLine = append(newProblemLine, numStr);
		} else if someRune == '\n' {
			problemInputMatrix = append(problemInputMatrix, newProblemLine);
			newProblemLine = []string{};
		} else if (someRune == '+') || (someRune == '*') {
			problemOperations = append(problemOperations, string(someRune))
			problemOperations = append(problemOperations, readOperations(reader)...);
		} else if err == io.EOF {
			break;
		}
	}

	problemInputMatrixInt := getCorrectNumsForProblems(problemInputMatrix);

	solutions := finishMathProblems(problemInputMatrixInt, problemOperations);
	total := 0
	for i := 0; i < len(solutions); i++ {
		total += solutions[i];
	}

	fmt.Println(total);
}
