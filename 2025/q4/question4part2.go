package main

import (
	"fmt"
	"os"
	"bufio"
)

type Point struct {
    x int
    y int
}

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

func buildMatrix(rows []string) [][]string {

	output := [][]string{}
	for i := 0; i < len(rows); i++ {
		curr := []string{}
		for j := 0; j < len(rows[i]); j++ {
			curr = append(curr, string(rows[i][j]))
		}
		output = append(output, curr)
	}

	return output
}

func inBound(x int, y int, xBound int, yBound int) bool {
	if x < 0 || x >= xBound {
		return false;
	}

	if y < 0 || y >= yBound {
		return false;
	}

	return true;
}

func checkAdjacentPositions(matrix [][]string, x int, y int, xBound int, yBound int) bool {
	count := 0;
	roll := "@";
	if (matrix[x][y] != roll) {
		return false;
	}
	for i := x - 1; i <= x + 1; i++ {
		for j := y - 1; j <= y + 1; j++ {
			if !inBound(i, j, xBound, yBound) {
				continue;
			}
			if i == x && j == y {
				continue;
			}
			if (matrix[i][j] == roll) {
				count++;
			}
		}
	}

	return count < 4;
}

var canMove = make(map[Point]struct{});

func runThroughMatrix(matrix [][]string, seen map[Point]struct{}, x int, y int, xBound int, yBound int, count int) int {
	if !inBound(x, y, xBound, yBound) {
		return count;
	}
	_, exists := seen[Point{x: x, y: y}];
	if exists {
		return count;
	}

	seen[Point{x: x, y: y}] = struct{}{};
	if checkAdjacentPositions(matrix, x, y, xBound, yBound) {
		count++;
		canMove[Point{x: x, y: y}] = struct{}{};
	}

	return max(runThroughMatrix(matrix, seen, x + 1, y, xBound, yBound, count),
		   runThroughMatrix(matrix, seen, x - 1, y, xBound, yBound, count),
		   runThroughMatrix(matrix, seen, x, y + 1, xBound, yBound, count),
		   runThroughMatrix(matrix, seen, x, y - 1, xBound, yBound, count))
}

func main() {
	file := "question4.input"
	rows, _ := readFile(file)
	matrix := buildMatrix(rows)
	total := 0
	for true {
		result := runThroughMatrix(matrix, make(map[Point]struct{}), 0, 0, len(matrix), len(matrix[0]), 0);
		total += result
		if result == 0 {
			break;
		}
		for point, _ := range canMove {
			matrix[point.x][point.y] = ".";
		}
	}
	fmt.Println(total);
}
