package main

import (
    "fmt"
	"strings"
	"strconv"
	
    "advent-of-code/pkg/utils"
)

// turnDial calculates the final dial position and how many times it passed zero.
func turnDial(start, dist int) (pos, rotations int) {
	end := start + dist

	rotations = utils.AbsInt(end / 100)

	if start != 0 && end <= 0 {
		rotations++
	}

	pos = utils.PMod(end, 100)
	return
}

// Part 1: Count the number of times the dial points to zero after a turn
// Part 2: Count the number of times the dial points to zero at any time during a turn
func main() {
	input := utils.ReadInput(1)
	lines := strings.Fields(input)

	dial := 50
	var p1, p2 int

	for _, line := range lines {
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(fmt.Errorf("invalid input %q: %w", line, err))
		}

		if dir == 'L' {
			dist = -dist
		}

		pos, rotations := turnDial(dial, dist)
		dial = pos

		if pos == 0 {
			p1++
		}
		p2 += rotations
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
