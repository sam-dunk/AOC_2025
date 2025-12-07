package main

import (
	"fmt"
	"strings"
	"slices"

	"advent-of-code/pkg/utils"
)

func main() {
	input := utils.ReadInput(7)
	lines := strings.Fields(input)

	r := len(lines)

	manifold := make([][] string, r)
	for i, line := range lines {
		manifold[i] = strings.Split(line, "")
	}

	sidx := slices.IndexFunc(manifold[0], func(s string) bool {
		return s == "S"
	})

	current := make([]int, len(manifold[0]))
	next := make([]int, len(manifold[0]))
	current[sidx] = 1

	p1 := 0

	for _, m := range manifold[1:] {
		clear(next)

		for j, v := range current {
			if v == 0 {
				continue
			}

			if m[j] == "^" {
				next[j-1] += v
				next[j+1] += v
				next[j] = 0
				
				p1++ 
			} else {
				next[j] += v
			}
		}

		current, next = next, current
	}

	p2 := 0
	for _, v := range current {
		p2 += v
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
