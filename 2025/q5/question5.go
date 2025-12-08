package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
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

func isInRange(ranges []string, id int) bool {
	for i := 0; i < len(ranges); i++ {
		rangeSplit := strings.Split(ranges[i], "-");
		lower, err := strconv.Atoi(rangeSplit[0]);
		if err != nil {
			panic(err);
		}
		upper, err := strconv.Atoi(rangeSplit[1]);
		if err != nil {
			panic(err);
		}
		if (id >= lower && id <= upper) {
			return true;
		}
	}

	return false;
}
func main() {
	lines, _ := readFile("question5.input");

	ranges := []string{};
	gatheringRanges := true;
	totalFresh := 0;

	for _, line := range lines {
		if gatheringRanges {
			if line == "" {
				gatheringRanges = false
			} else {
				ranges = append(ranges, line) 
			}
		} else {
			id, err := strconv.Atoi(line);
			if err != nil {
				panic(err);
			}
			if isInRange(ranges, id) {
				totalFresh++;
			}
		}
	}

	fmt.Println(totalFresh);
}