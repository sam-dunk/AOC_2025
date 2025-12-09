package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"

	"advent-of-code/pkg/utils"
)

type Tile struct {
	X int
	Y int
}

func tileArea(t1, t2 Tile) (a int) {
	a = (utils.AbsInt(t2.X - t1.X) + 1) * (utils.AbsInt(t2.Y - t1.Y) + 1)
	return
}

type Rectangle struct {
	Area int
	Point1 Tile
	Point2 Tile
}

func isIntersecting(r Rectangle, t1, t2 Tile) bool {
	left := min(r.Point1.X, r.Point2.X)
	right := max(r.Point1.X, r.Point2.X)
	top := min(r.Point1.Y, r.Point2.Y)
	bot := max(r.Point1.Y, r.Point2.Y)

	if t1.X == t2.X {
		x := t1.X
		if x > left && x < right {
			minY := min(t1.Y, t2.Y)
			maxY := max(t1.Y, t2.Y)
			if minY <= top && maxY > top {
				return true
			}
			if minY < bot && maxY >= bot {
				return true
			}
		}
	} else {
		y := t1.Y
		if y > top && y < bot {
			minX := min(t1.X, t2.X)
			maxX := max(t1.X, t2.X)
			if minX <= left && maxX > left {
				return true
			}
			if minX < right && maxX >= right {
				return true
			}
		}
	}

	return false
}

func main() {
	input := utils.ReadInput(9)
	lines := strings.Fields(input)

	redTiles := make([]Tile, len(lines))

	for i, line := range lines {
		sline := strings.Split(line, ",")
		x, _ := strconv.Atoi(sline[0])
		y, _ := strconv.Atoi(sline[1])
		
		redTiles[i] = Tile{x, y}
	}

	rectangles := []Rectangle{}
	for i, tile1 := range redTiles {
		for _, tile2 := range redTiles[i+1:] {
			rectangles = append(rectangles, Rectangle{tileArea(tile1, tile2), tile1, tile2})
		}
	}

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Area > rectangles[j].Area
	})

	//
	// At this point, the rectangles slice contains every unique rectangle we can make from pairs of red tiles.
	// This slice is sorted in decreasing order of total area, meaning the solution to part 1 (the largest such
	// rectangle) is just the first entry in the slice.
	//

	fmt.Println("Part 1:", rectangles[0].Area)

	//
	// For part 2, we can iterate over this slice until we find the first rectangle completely contained within 
	// the polygon formed by redTiles.
	//
	// Note: To determine if a rectangle is completely contained, I am checking whether the path ever crosses
	// from the the border/outside to the inside of the rectangle. It is possible that the rectangle's interior
	// falls completely outside the polygon, however, this was not an issue with my input. E.g.,
	//
	//										# # # # # # # # # # # # # 
	//										#     X # # # # # #     #
 	//										#     #           #     #  
	//										#     #           #     #  
	//										# # # #           X # # #
	//
	// Here, we would incorrectly select the rectangle formed by the two X's as our solution, despite the fact it 
	// clearly is NOT contained in the polygon's interior
	//		
	
	for _, r := range rectangles {
		contained := true

		for i, p2 := range redTiles {
			var p1 Tile
			if i == 0 {
				p1 = redTiles[len(redTiles)-1]
			} else {
				p1 = redTiles[i-1]
			}

			if isIntersecting(r, p1, p2) {
				contained = false 
				break
			}
		}

		if contained {
			fmt.Println(r.Area)
			break
		}
	}
}