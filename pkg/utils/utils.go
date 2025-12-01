package utils

import (
    "fmt"
    "os"
)

func ReadInput(day int) string {
    data, err := os.ReadFile(fmt.Sprintf("inputs/day%02d.txt", day))
    if err != nil {
        panic(err)
    }
    return string(data)
}

// Positive modulo, returns non negative solution to x % d  
func PMod(x, d int) int {  
  x = x % d
  if x >= 0 { return x }
  if d < 0 { return x - d }
  return x + d
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}