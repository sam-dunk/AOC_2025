package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"advent-of-code/pkg/utils"
)

type IdRange struct {
	Lower int
	Upper int
}

func main() {
	input := utils.ReadInput(5)
	lines := strings.Split(input, "\n")
	idx := slices.IndexFunc(lines, func(s string) bool {
		return s == ""
	})

	ranges := lines[:idx]
	ids := lines[idx+1:]

	rangeList := make([]IdRange, len(ranges))
	for i, r := range ranges {
		splitR := strings.Split(r, "-")
		l, _ := strconv.Atoi(splitR[0])
		h, _ := strconv.Atoi(splitR[1])
		rangeList[i] = IdRange{l, h}
	}

	sort.Slice(rangeList, func(i, j int) bool {
		return rangeList[i].Lower < rangeList[j].Lower
	})

	intIDs := make([]int, len(ids))
	for i, id := range ids {
		intIDs[i], _ = strconv.Atoi(id)
	}

	// Part 1: Count how many IDs fall within any range
	p1 := 0
	for _, id := range intIDs {
		for _, pair := range rangeList {
			if id < pair.Lower {
				break // Ranges are sorted, no point checking further
			}
			if id <= pair.Upper {
				p1++
				break
			}
		}
	}

	// Part 2: Count total unique IDs covered by all ranges (merging overlaps)
	p2 := 0
	bound := -1 // Tracks the highest ID we've counted so far
	for _, r := range rangeList {
		l := r.Lower
		h := r.Upper

		// Skip ranges completely contained within already-counted lines
		if h <= bound {
			continue
		}

		// Adjust start if this range overlaps with an already-counted line
		if l <= bound {
			l = bound + 1
		}

		p2 += h - l + 1
		bound = h
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}