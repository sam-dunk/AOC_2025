package main

import (
    "fmt"
	"strings"
	"strconv"

    "advent-of-code/pkg/utils"
)

func main() {
	input := utils.ReadInput(2)
	lines := strings.Split(input, ",")

	p1 := 0
	p2 := 0

	for _, line := range lines {
		r := strings.Split(line, "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])

		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)

			mid := len(id) / 2
			if id[:mid] == id[mid:] {
				p1 += i
				p2 += i
				continue
			}

			for j := 1; j <= len(id) / 2; j++ {
				p := id[:j]

				if len(id) % len(p) != 0 {
					continue
				}

				if strings.Count(id, p) == len(id) / len(p) {
					p2 += i 
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
