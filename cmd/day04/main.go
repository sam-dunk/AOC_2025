package main

import (
    "fmt"
	"strings"

    "advent-of-code/pkg/utils"
)

func is_accessible(matrix [][]int, i, j int) bool {
	neighbors := 0

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}

			if matrix[i + x][j + y] == 1 {
				neighbors++
			}
		}
	}

	return neighbors < 4
}

func main() {
	input := utils.ReadInput(4)
	lines := strings.Fields(input)

	rows := len(lines)
	cols := len(lines[0])

	matrix := make([][]int, rows + 2)
	for i := range matrix {
		matrix[i] = make([]int, cols + 2)
	}

	for i, line := range lines {
		for j, char := range line {
			if char == '@' {
				matrix[i+1][j+1] = 1
			}
		}
	}

	p1 := 0
	p2 := 0

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if is_accessible(matrix, i, j) && matrix[i][j] == 1 {
				p1++
			}
		}
	}

	has_changed := true
	for has_changed {
		has_changed = false
		for i := 1; i <= rows; i++ {
			for j := 1; j <= cols; j++ {
				if is_accessible(matrix, i, j) && matrix[i][j] == 1 {
					matrix[i][j] = 0
					p2++
					has_changed = true
				}
			}
		}
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
