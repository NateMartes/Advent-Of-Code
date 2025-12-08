package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
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

func getBounds(someRange string) (int, int) {
	rangeSplit := strings.Split(someRange, "-");
	lower, err := strconv.Atoi(rangeSplit[0]);
	if err != nil {
		panic(err);
	}
	upper, err := strconv.Atoi(rangeSplit[1]);
	if err != nil {
		panic(err);
	}

	return lower, upper;
}

func getRangeSize(someRange string) int {
	rangeSplit := strings.Split(someRange, "-");
	lower, err := strconv.Atoi(rangeSplit[0]);
	if err != nil {
		panic(err);
	}
	upper, err := strconv.Atoi(rangeSplit[1]);
	if err != nil {
		panic(err);
	}

	return (upper - lower) + 1;
}

func addRange(someRange string, ranges []string) []string {
	ranges = append(ranges, someRange)
	sort.Slice(ranges, func(i, j int) bool {

		leftLower, _ := strconv.Atoi(strings.Split(ranges[i], "-")[0]);
		leftUpper, _ := strconv.Atoi(strings.Split(ranges[i], "-")[1]);
		rightLower, _ := strconv.Atoi(strings.Split(ranges[j], "-")[0]);
		rightUpper, _ := strconv.Atoi(strings.Split(ranges[j], "-")[1]);
		
		if leftLower != rightLower {
			return leftLower < rightLower;
		}

		return leftUpper < rightUpper;
	});

	return ranges;
}

func fixRanges(ranges []string) []string {
	output := []string{};
	i := 0;
 	for i < len(ranges) {
		j := i + 1;
		if j == len(ranges) { output = append(output, ranges[i]); break; };

		leftLower, _ := strconv.Atoi(strings.Split(ranges[i], "-")[0]);
		leftUpper, _ := strconv.Atoi(strings.Split(ranges[i], "-")[1]);
		rightLower, _ := strconv.Atoi(strings.Split(ranges[j], "-")[0]);
		rightUpper, _ := strconv.Atoi(strings.Split(ranges[j], "-")[1]);

		if (rightLower <= leftUpper) && (rightUpper >= leftLower) && (rightUpper >= leftUpper) {
			fmt.Printf("Merging left range %s with right range %s\n", ranges[i], ranges[j]);
			output = append(output, strconv.Itoa(leftLower) + "-" + strconv.Itoa(rightUpper));
			i += 2;
		} else if (leftUpper >= rightLower) && (leftUpper >= rightUpper) {
			fmt.Printf("Overpower left range %s on right range %s\n", ranges[i], ranges[j]);
			output = append(output, ranges[i]);
			i += 2;
		} else {
			fmt.Printf("No comparision from left range %s and right range %s\n", ranges[i], ranges[j]);
			output = append(output, ranges[i]);
			i++;
		}
	}
	fmt.Println(output);
	return output;
}

func sizeDiff(a []string, b []string) bool {
	return len(a) != len(b);
}

func main() {

	lines, _ := readFile("question5.input");

	ranges := []string{}
	for _, line := range lines {
		if line == "" {
			break;
		} else {
			if len(ranges) == 0 {
				ranges = append(ranges, line);
			} else {
				ranges = addRange(line, ranges);
				fmt.Println(ranges);
				tmp := fixRanges(ranges);
				for sizeDiff(ranges, tmp) {
					ranges = tmp;
					tmp = fixRanges(ranges);
				}
			}
		}
	}

	total := 0
	for _, someRange := range ranges {
		total += getRangeSize(someRange);
	}
	fmt.Println(total);
}