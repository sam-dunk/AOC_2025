package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/pkg/utils"
)

func main() {
	input := utils.ReadInput(6)
	lines := strings.Split(input, "\n")
	
	nums1 := strings.Fields(lines[0])
	nums2 := strings.Fields(lines[1])
	nums3 := strings.Fields(lines[2])
	nums4 := strings.Fields(lines[3])
	ops := strings.Fields(lines[4])
	results := make([]int, len(ops))

	p1 := 0
	p2 := 0

	for i := 0; i < len(ops); i++ {
		int1, _ := strconv.Atoi(nums1[i])
		int2, _ := strconv.Atoi(nums2[i])
		int3, _ := strconv.Atoi(nums3[i])
		int4, _ := strconv.Atoi(nums4[i])

		if ops[i] == "+" {
			results[i] = int1 + int2 + int3 + int4
			p1 += results[i]
		} else {
			results[i] = int1 * int2 * int3 * int4
			p1 += results[i]
		}
	}

	nums1 = strings.Split(lines[0], "")
	nums2 = strings.Split(lines[1], "")
	nums3 = strings.Split(lines[2], "")
	nums4 = strings.Split(lines[3], "")
	ops = strings.Split(lines[4], "")

	last_op := ""
	acc := 0
	for i := 0; i < len(ops); i++ {
		if nums1[i] == " " && nums2[i] == " " && nums3[i] == " " && nums4[i] == " " {
			p2 += acc
			continue
		}

		if ops[i] != " " {
			last_op = ops[i]
			if last_op == "+" {
				acc = 0
			} else {
				acc = 1
			}
		}

		curr_num := 0

		if nums1[i] != " " {
			n1, _ := strconv.Atoi(nums1[i])
			curr_num += n1
		}

		if nums2[i] != " " {
			n2, _ := strconv.Atoi(nums2[i])
			curr_num = curr_num * 10 + n2
		}

		if nums3[i] != " " {
			n3, _ := strconv.Atoi(nums3[i])
			curr_num = curr_num * 10 + n3
		}

		if nums4[i] != " " {
			n4, _ := strconv.Atoi(nums4[i])
			curr_num = curr_num * 10 + n4
		}

		if last_op == "+" {
			acc += curr_num
		} else {
			acc *= curr_num
		}

		if i + 1 >= len(ops) {
			p2 += acc
		}

	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}