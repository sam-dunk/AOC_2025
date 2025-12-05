package main

import (
    "fmt"
	"strings"
	"strconv"

    "advent-of-code/pkg/utils"
)

func get_largest(str string) (idx, largest int) {
	largest_str := str[:1]
	idx = 0
	for i := 1; i < len(str); i++ {
		if str[i:i+1] > largest_str {
			largest_str = str[i:i+1]
			idx = i
		}
	}

	largest, _ = strconv.Atoi(largest_str)
	return
}

func main() {
	input := utils.ReadInput(3)
	banks := strings.Split(input, "\n")

	p1 := 0
	p2 := 0

	for _, bank := range banks {
		l_idx, first_digit := get_largest(bank[:len(bank)-1])
		_, second_digit := get_largest(bank[l_idx + 1:])

		p1 += first_digit * 10 + second_digit
	}

	for _, bank := range banks {
		acc := 0
		search_start := 0
		
		// Find 12 largest digits in order
		for digits_found := 0; digits_found < 12; digits_found++ {
			digits_remaining := 12 - digits_found
			// Only search up to the point where we can still get enough digits
			search_end := len(bank) - digits_remaining + 1
			
			// Search for largest digit in valid range
			relative_idx, n_digit := get_largest(bank[search_start:search_end])
			
			// Update absolute position for next search
			search_start = search_start + relative_idx + 1
			
			acc = acc * 10 + n_digit 
		}
		p2 += acc
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
