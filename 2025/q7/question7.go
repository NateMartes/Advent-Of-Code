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

func isBeam(matrix *[][]rune, x int, y int, xBound int, yBound int) bool {
	if (x < 0) || (x > xBound) {
		return false;
	}
	if (y < 0) || (y > yBound) {
		return false;
	}

	if ((*matrix)[y][x] == '|') {
		return true;
	} else {
		return false;
	}
}

func addBeams(matrix *[][]rune, startX int, startY int, xBound int, yBound int) {
	if (startX < 0) || (startX > xBound) {
		return;
	}
	if (startY < 0) || (startY > yBound) {
		return;
	}

	for (startX < xBound) && (startY < yBound) {
		someRune := (*matrix)[startY][startX];
		if someRune == '^' {
			return;
		} else {
			(*matrix)[startY][startX] = '|';
		}
		startY++;
	}

	return;
}

func getSplitCount(matrix *[][]rune, startX int, startY int, xBound int, yBound int) int {
	if (startX < 0) || (startX > xBound) {
		return 0;
	}
	if (startY < 0) || (startY > yBound) {
		return 0;
	}
	count := 0;

	splitterX, splitterY, splitterFound := findNextSpiltter(matrix, startX, startY, xBound, yBound);
	if splitterFound {

		if !isBeam(matrix, splitterX - 1, splitterY, xBound, yBound) || !isBeam(matrix, splitterX + 1, splitterY, xBound, yBound) {
			count++;
		}

		if !isSplitter(matrix, splitterX + 1, splitterY, xBound, yBound) && !isBeam(matrix, splitterX + 1, splitterY, xBound, yBound) {
			(*matrix)[splitterY][splitterX + 1] = '|';
			addBeams(matrix, splitterX + 1, splitterY, xBound, yBound);
			count += getSplitCount(matrix, splitterX + 1, splitterY, xBound, yBound);
		}

		if !isSplitter(matrix, splitterX - 1, splitterY, xBound, yBound) && !isBeam(matrix, splitterX - 1, splitterY, xBound, yBound) {
			(*matrix)[splitterY][splitterX - 1] = '|';
			addBeams(matrix, splitterX - 1, splitterY, xBound, yBound);
			count += getSplitCount(matrix, splitterX - 1, splitterY, xBound, yBound);
		}

		return count;

	} else {
		return count;
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

	fmt.Println(getSplitCount(&matrix, startX, startY, len(matrix[0]), len(matrix)));
	fmt.Println("Matrix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println()
	}
}
