package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
)

func readFile(name string) (string, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	str := string(b)
    return str, nil
}

type xy struct {
    x int
    y int
}

func findNextSpiltter(matrix *[][]rune, startX int, startY int, xBound int, yBound int) (int, int, bool) {
	if (startX < 0) || (startX > xBound) {
		return -1, -1, false;
	}

	if (startY < 0) || (startY > yBound) {
		return -1, -1, false;
	}

	for (startX < xBound) && (startY < yBound) {
		someRune := (*matrix)[startY][startX];
		if someRune == '^' {
			return startX, startY, true;
		}
		startY++;
	}

	return -1, -1, false;
}

func isSplitter(matrix *[][]rune, x int, y int, xBound int, yBound int) bool {
	if (x < 0) || (x > xBound) {
		return false;
	}
	if (y < 0) || (y > yBound) {
		return false;
	}

	if ((*matrix)[y][x] == '^') {
		return true;
	} else {
		return false;
	}
}

func getSplitCount(matrix *[][]rune, startX int, startY int, xBound int, yBound int, cache map[xy]int) int {
	
	if (startX < 0) || (startX > xBound) {
		return 0;
	}

	if (startY < 0) || (startY > yBound) {
		return 0;
	}

    key := xy{startX, startY}
    if v, ok := cache[key]; ok {
        return v
    }


	fmt.Println(startX, startY);

	splitterX, splitterY, splitterFound := findNextSpiltter(matrix, startX, startY, xBound, yBound);

	if splitterFound {

		count := 0;

		if !isSplitter(matrix, splitterX + 1, splitterY, xBound, yBound) {
			count += getSplitCount(matrix, splitterX + 1, splitterY, xBound, yBound, cache);
		}

		if !isSplitter(matrix, splitterX - 1, splitterY, xBound, yBound) {
			count += getSplitCount(matrix, splitterX - 1, splitterY, xBound, yBound, cache);
		}

		cache[key] = count
		return count;

	} else {

		cache[key] = 1;
		return 1;
	}
}

func main() {

	fileStr, _ := readFile("question7.input");
	reader := bufio.NewReader(strings.NewReader(fileStr));

	startX := 0;
	startY := 0;

	x := 0;
	y := 0;

	matrix := [][]rune{};
	matrixLine := []rune{};

	for {

		someRune, _, err := reader.ReadRune();
		matrixLine = append(matrixLine, someRune);
		if someRune == 'S' {
			startX = x;
			startY = y;
		} else if someRune == '\n' {
			matrix = append(matrix, matrixLine);
			matrixLine = []rune{};
			y++;
		} else if err == io.EOF {
			break;
		}
		x++;

	}

	fmt.Println(getSplitCount(&matrix, startX, startY, len(matrix[0]), len(matrix),  make(map[xy]int)));
}