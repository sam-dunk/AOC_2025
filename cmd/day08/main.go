package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"sort"

	"advent-of-code/pkg/utils"
)

func dist(x1, y1, z1, x2, y2, z2 int) (dist float64) {
	dist = math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2) + math.Pow(float64(z2-z1),2)
	dist = math.Sqrt(dist)
	return
}

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func union(parent, size []int, i, j int) bool {
	rooti := find(parent, i)
	rootj := find(parent, j)
	
	if rooti == rootj {
		return false
	}
	
	if size[rooti] < size[rootj] {
		parent[rooti] = rootj
		size[rootj] += size[rooti]
	} else {
		parent[rootj] = rooti
		size[rooti] += size[rootj]
	}

	return true
}

func main() {
	input := utils.ReadInput(8)
	lines := strings.Fields(input)

	boxes := make([][]int, len(lines))

	for i, line := range lines {
		split := strings.Split(line, ",")
		ints := []int{}

		for _, s := range split {
			v, _ := strconv.Atoi(s)
			ints = append(ints, v)
		}

		boxes[i] = ints
	}

	dmat := make([][]float64, len(lines))

	for i := range boxes {
		dists := []float64{}

		for j := range boxes {
			if j == i {
				dists = append(dists, math.MaxFloat64)
			} else {
				dists = append(dists, dist(boxes[i][0], boxes[i][1], boxes[i][2], boxes[j][0], boxes[j][1], boxes[j][2]))
			}
		}
		dmat[i] = dists
	}

	parent := make([]int, len(boxes))
	size := make([]int, len(boxes))
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	// Part 1: Make 1000 connections
	for k := 0; k < 1000; k++ {
		minDist := math.MaxFloat64
		mini, minj := -1, -1
		
		for i := range dmat {
			for j := range dmat[i] {
				if dmat[i][j] < minDist {
					minDist = dmat[i][j]
					mini = i
					minj = j
				}
			}
		}
		
		dmat[mini][minj] = math.MaxFloat64
		dmat[minj][mini] = math.MaxFloat64
		
		union(parent, size, mini, minj)
	}

	circuitSizes := make(map[int]int)
	for i := range boxes {
		root := find(parent, i)
		circuitSizes[root] = size[root]
	}

	var circuits []int
	for _, s := range circuitSizes {
		circuits = append(circuits, s)
	}

	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i] > circuits[j]
	})
	
	p1 := circuits[0] * circuits[1] * circuits[2]

	// Part 2: Continue until all in one circuit
	numCircuits := len(circuitSizes)
	var lastI, lastJ int
	
	for numCircuits > 1 {
		minDist := math.MaxFloat64
		mini, minj := -1, -1
		
		for i := range dmat {
			for j := range dmat[i] {
				if dmat[i][j] < minDist {
					minDist = dmat[i][j]
					mini = i
					minj = j
				}
			}
		}
		
		dmat[mini][minj] = math.MaxFloat64
		dmat[minj][mini] = math.MaxFloat64
		
		if union(parent, size, mini, minj) {
			lastI = mini
			lastJ = minj
			numCircuits--
		}
	}

	p2 := boxes[lastI][0] * boxes[lastJ][0]

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}